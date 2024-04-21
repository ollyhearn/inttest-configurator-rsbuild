package projects

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type API struct {
	log     *zap.SugaredLogger
	useCase IUseCase
}

func New(useCase IUseCase, log *zap.SugaredLogger) *API {
	return &API{
		log:     log,
		useCase: useCase,
	}
}

func (a *API) Register(router fiber.Router, authMiddleware fiber.Handler, middlewares ...fiber.Handler) {

}
