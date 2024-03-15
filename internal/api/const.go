package api

import (
	"github.com/gofiber/fiber/v2"
)

const JwtLocation = "header:" + fiber.HeaderAuthorization
const JwtSecret = "mysecret"
