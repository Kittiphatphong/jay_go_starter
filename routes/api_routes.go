package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/controllers/api"
)

type apiRoutes struct {
	controllerApi api.ControllerApi
}

func (a apiRoutes) Install(app *fiber.App) {
	route := app.Group("api/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Post("hello", a.controllerApi.StartController)
}

func NewApiRoutes(
	controllerApi api.ControllerApi,
	//controller
) Routes {
	return &apiRoutes{
		controllerApi: controllerApi,
		//controller
	}
}
