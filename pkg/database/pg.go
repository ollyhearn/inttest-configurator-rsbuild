package database

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"go.uber.org/zap"
)

type PGDB struct {
	log *zap.SugaredLogger
	*pg.DB
}

func (db *PGDB) Pg() *pg.DB {
	return db.DB
}

func NewPostgres(
	ctx context.Context,
	connUrl string,
	log *zap.SugaredLogger,
	additionalOpts ...PostgresInitOpt,
) (*PGDB, error) {
	opt, err := pg.ParseURL(connUrl)
	if err != nil {
		return nil, err
	}
	for _, o := range additionalOpts {
		o(opt)
	}
	db := &PGDB{
		DB:  pg.Connect(opt).WithContext(ctx),
		log: log,
	}
	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}
	db.AddQueryHook(&pgQueryLogger{log: log.With(zap.Namespace("PGDB"))})
	return db, nil
}

type PostgresInitOpt func(opts *pg.Options)

func WithPgPoolSize(poolSize int) PostgresInitOpt {
	return func(opts *pg.Options) {
		opts.PoolSize = poolSize
	}
}

type pgQueryLogger struct {
	log *zap.SugaredLogger
}

// AfterQuery implements pg.QueryHook.
func (q *pgQueryLogger) AfterQuery(ctx context.Context, event *pg.QueryEvent) error {
	q.log.Debugf("query executed in %d ms", time.Since(event.StartTime).Milliseconds())
	return nil
}

// BeforeQuery implements pg.QueryHook.
func (q *pgQueryLogger) BeforeQuery(ctx context.Context, event *pg.QueryEvent) (context.Context, error) {
	query, err := event.FormattedQuery()
	if err != nil {
		panic(err)
	}

	log := q.log

	switch event.DB.(type) {
	case *pg.DB:
		log = log.With(zap.Bool("DB", true))
	case *pg.Tx:
		log = log.With(zap.Bool("TX", true))
	}
	log.Debug(string(query))
	return ctx, nil
}

var _ pg.QueryHook = (*pgQueryLogger)(nil)

func (db *PGDB) Ping(ctx context.Context) error {
	var ping int
	_, err := db.WithContext(ctx).QueryOne(pg.Scan(&ping), "SELECT 1")
	return err
}

func (db *PGDB) RunInTransaction(ctx context.Context, fn func(db orm.DB) error) error {
	if tx, ok := GetPgTx(ctx); ok {
		return fn(tx)
	}
	return db.DB.RunInTransaction(ctx, func(tx *pg.Tx) error {
		return fn(tx)
	})
}

func (db *PGDB) TxContext(ctx context.Context) (context.Context, error) {
	if _, ok := GetPgTx(ctx); !ok {
		tx, err := db.BeginContext(ctx)
		if err != nil {
			return nil, err
		}
		return context.WithValue(ctx, ctxPgTxKey, tx), nil
	}
	return ctx, nil
}

type pgCtxKey string

const (
	ctxPgTxKey pgCtxKey = "internalctx_postgres_tx"
)

func GetPgTx(ctx context.Context) (*pg.Tx, bool) {
	v := ctx.Value(ctxPgTxKey)
	if v == nil {
		return nil, false
	}
	tx, ok := v.(*pg.Tx)
	if !ok {
		panic("runtime error: excepted *PGTX type in transaction extraction")
	}
	return tx, true
}
