package http_handler

import (
	"fmt"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-time/api"
	gen "go-time/generated"
)

func SetupWebServer() {

	prom := fiberprometheus.New("test")
	server := api.NewServer()

	app := fiber.New()

	prom.RegisterAt(app, "/metrics")
	app.Use(prom.Middleware)
	gen.RegisterHandlers(app, server)

	app.Static("/swagger/api/openapi.yaml", "./api/openapi.yaml")
	app.Get("/swagger/*", adaptor.HTTPHandlerFunc(
		httpSwagger.Handler(

			httpSwagger.URL("/swagger/api/openapi.yaml"),
		),
	))

	err := app.Listen(":3000")

	if err != nil {
		fmt.Println("Unable to start server", err)
	}
}
