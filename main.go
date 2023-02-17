package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go_starter/config"
	"go_starter/controllers"
	"go_starter/controllers/api"
	"go_starter/database"
	"go_starter/logs"
	"go_starter/partners"
	"go_starter/repositories"
	"go_starter/routes"
	"go_starter/services"
	"go_starter/trails"
	"log"
	"net/http"
)

func main() {

	//connect database
	postgresConnection, err := database.PostgresConnection()
	if err != nil {
		logs.Error(err)
		return
	}

	//call api client interface
	httpClient := http.Client{}
	newHttpClientTrail := trails.NewHttpClientTrail(httpClient)
	partners.NewPartner(newHttpClientTrail)

	//basic structure
	newRepository := repositories.NewRepository(postgresConnection)
	newService := services.NewService(newRepository)

	//connect route
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New())
	app.Use(cors.New())

	//Web routes
	newController := controllers.NewController(newService)
	newWebRoute := routes.NewWebRoutes(
		newController,
		//new web controller
	)
	newWebRoute.Install(app)

	//Api routes
	newControllerApi := api.NewControllerApi(newService)
	newApiRoute := routes.NewApiRoutes(
		newControllerApi,
		//new web controller
	)
	newApiRoute.Install(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env("app.port"))))
}
