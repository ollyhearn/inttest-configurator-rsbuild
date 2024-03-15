package auth

import (
	"errors"

	"github.com/Ghytro/inttest-configurator/internal/api"
	"github.com/Ghytro/inttest-configurator/internal/entity"
	entAuth "github.com/Ghytro/inttest-configurator/internal/entity/auth"
	"github.com/Ghytro/inttest-configurator/internal/usecase/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type API struct {
	log     *zap.SugaredLogger
	useCase *auth.UseCase // пока без интерфейса обойдусь
}

func New(log *zap.SugaredLogger, useCase *auth.UseCase) *API {
	return &API{
		log:     log.With(zap.String("api", "auth_api")),
		useCase: useCase,
	}
}

func (a *API) Register(router fiber.Router, authMiddleware fiber.Handler, middlewares ...fiber.Handler) {
	r := fiber.New()
	r.Post("/", a.auth)
	router.Mount("/auth", r)

	r = fiber.New()

	r.Use(authMiddleware)
	for _, m := range middlewares {
		r.Use(m)
	}

	r.Get("/", a.listUsers)
	r.Post("/", a.createUser)
	r.Delete("/:id", a.deleteUser)

	router.Mount("/users", r)
}

func (a *API) createUser(ctx *fiber.Ctx) error {
	const location = "ошибка создания пользователя"

	form, err := api.ParseBody[createUserRequest](ctx)
	if err != nil {
		return err
	}
	creator, ok := api.GetUserEntity(ctx)
	if !ok {
		return errors.New("user not authorized")
	}
	result, err := a.useCase.CreateUser(ctx.Context(), creator.Id, form.UserName, form.Password, form.Roles...)
	if err != nil {
		return err
	}
	return ctx.JSON(createUserResponse{
		Id:        result.Id,
		CreatedAt: result.CreatedAt,
	})
}

func (a *API) deleteUser(ctx *fiber.Ctx) error {
	const location = "ошибка удаления пользователя"

	id, err := api.ParseUrlParamsId[entity.BigIntPK](ctx, "id")
	if err != nil {
		return err
	}
	deleter, ok := api.GetUserEntity(ctx)
	if !ok {
		return errors.New("user not authorized")
	}
	err = a.useCase.DeleteUser(ctx.Context(), deleter.Id, id)
	if err != nil {
		return err
	}
	return ctx.JSON(nil)
}

func (a *API) listUsers(ctx *fiber.Ctx) error {
	const location = "ошибка получения списка пользователей"

	querier, ok := api.GetUserEntity(ctx)
	if !ok {
		return errors.New("no user entity in authorized ctx")
	}
	users, err := a.useCase.ListUsers(ctx.Context(), querier.Id)
	if err != nil {
		return err
	}
	return ctx.JSON(lo.Map(users, func(u entAuth.User, _ int) listUsersResponseItem {
		return listUsersResponseItem{
			Id:        u.Id,
			UserName:  u.UserName,
			CreatedAt: u.CreatedAt,
			Roles: lo.Map(u.Roles, func(r entAuth.Role, _ int) string {
				return r.Name
			}),
		}
	}))
}

func (a *API) auth(ctx *fiber.Ctx) error {
	form, err := api.ParseBody[authRequest](ctx)
	if err != nil {
		return err
	}
	token, err := a.useCase.GenToken(ctx.Context(), form.UserName, form.Password)
	if err != nil {
		return err
	}
	return ctx.JSON(authResponse{
		Token: token,
	})
}
