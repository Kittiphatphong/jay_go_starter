package api

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/controllers"
	"go_starter/services"
)

type ControllerApi interface {
	StartController(ctx *fiber.Ctx) error
}

type controllerApi struct {
	service services.Service
}

func (c controllerApi) StartController(ctx *fiber.Ctx) error {
	//TODO implement me
	return controllers.NewSuccessMsg(ctx, "Hello Golang app api")
}

func NewControllerApi(
	service services.Service,
	//service
) ControllerApi {
	return &controllerApi{
		service: service,
		//service
	}
}
