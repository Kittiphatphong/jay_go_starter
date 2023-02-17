package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/controllers"
)

type webRoutes struct {
	controller controllers.Controller
}

func (w webRoutes) Install(app *fiber.App) {
	route := app.Group("web/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Post("hello", w.controller.StartController)

}

func NewWebRoutes(
	controller controllers.Controller,
	//controller
) Routes {
	return &webRoutes{
		controller: controller,
		//controller
	}
}
