package api

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Register(router fiber.Router, authMiddleware fiber.Handler, middlewares ...fiber.Handler)
}
