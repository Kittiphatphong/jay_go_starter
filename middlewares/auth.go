package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"go_starter/config"
	"net/http"
)

func NewAuthWebMiddleware(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Env("auth.secretWeb")),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status": http.StatusUnauthorized,
				"error":  "UNAUTHORIZED",
			})
		},
	})(ctx)

}

func NewAuthApiMiddleware(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Env("auth.secretApi")),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status": http.StatusUnauthorized,
				"error":  "UNAUTHORIZED",
			})
		},
	})(ctx)
}