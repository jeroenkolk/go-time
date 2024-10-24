package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-time/internal/handlers"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.Hello)
	app.Get("/time", handlers.GetTime)
}
