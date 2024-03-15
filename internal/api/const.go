package api

import (
	"github.com/gofiber/fiber/v2"
)

const JwtLocation = "header:" + fiber.HeaderAuthorization
const JwtContextKey = "jwt_user_fiber_context"
const UserEntityContextKey = "__user_entity"
