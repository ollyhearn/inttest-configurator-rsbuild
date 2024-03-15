package api

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var AuthMiddleware = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{
		Key: []byte(JwtSecret),
	},
	TokenLookup: JwtLocation,
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if err.Error() == "Missing or malformed JWT" {
			// ErrResponse
			return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует или неправильно сформирован Токен Авторизации")
		} else {
			// ErrResponse
			return c.Status(fiber.StatusUnauthorized).SendString("Недействительный или просроченный Токен Авторизации")
		}
	},
})

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	return err
}
