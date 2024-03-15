package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"

	_ "github.com/lib/pq"

	"github.com/Ghytro/inttest-configurator/internal/api"
	authApi "github.com/Ghytro/inttest-configurator/internal/api/auth"
	authRepository "github.com/Ghytro/inttest-configurator/internal/repository/auth"
	authUseCase "github.com/Ghytro/inttest-configurator/internal/usecase/auth"
	"github.com/Ghytro/inttest-configurator/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/pressly/goose"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

func main() {
	// fixme: по нормальному через кобру переделать
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	sigIntListener(g)

	logger, err := initLogger()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error in logger init"))
	}

	db, err := database.NewPostgres(
		ctx,
		os.Getenv("DATABASE_URL"), // separate the config
		logger,
		database.WithPgPoolSize(8),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error connecting to db"))
	}
	if os.Getenv("AUTO_MIGRATE") == "true" {
		if err := autoMigrate(os.Getenv("DATABASE_URL"), "up"); err != nil {
			log.Fatal(errors.Wrap(err, "error while migrating"))
		}
	}

	apis := initApis(logger, db)
	app := initFiberRouter(apis)
	ln, err := initFiberListener(":8080")
	if err != nil {
		log.Fatal(err)
	}
	listenFiber(ctx, g, app, ln)
}

func sigIntListener(g *errgroup.Group) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	g.Go(func() error {
		for range sigs {
			return errors.New("sigint caught, graceful shutdown")
		}
		return nil
	})
}

func initLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewDevelopment(zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

func initApis(logger *zap.SugaredLogger, db *database.PGDB) (apis []api.Handler) {
	{
		repo := authRepository.New(db, logger)
		useCase := authUseCase.New(logger, repo)
		apis = append(apis, authApi.New(logger, useCase))
	}
	return
}

func initFiberRouter(apis []api.Handler) *fiber.App {
	fiberApi := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		ErrorHandler: api.ErrorHandler,

		BodyLimit:                100 * 1024 * 1024, // 300mb
		EnableSplittingOnParsers: true,
	})
	apiV1 := fiberApi.Group("/api/v1")
	for _, a := range apis {
		a.Register(apiV1, api.AuthMiddleware)
	}
	return fiberApi
}

func initFiberListener(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}

func listenFiber(ctx context.Context, g *errgroup.Group, app *fiber.App, ln net.Listener) {
	g.Go(func() error {
		return app.Listener(ln)
	})
	<-ctx.Done()
	ln.Close()
}

const (
	pgDialect       = "postgres"
	sqlMigrationDir = "./internal/migrations"
)

func autoMigrate(dbUrl, cmd string) error {
	if err := goose.SetDialect(pgDialect); err != nil {
		return err
	}

	db, err := sql.Open(pgDialect, dbUrl)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic("failed to close goose dbconn")
		}
	}()
	if err := goose.Run(cmd, db, sqlMigrationDir); err != nil {
		return err
	}
	return nil
}
