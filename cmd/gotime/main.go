package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-time/internal/routes"
)

func main() {
	app := fiber.New()
	routes.Setup(app)
	err := app.Listen(":3000")

	if err != nil {
		fmt.Println("Unable to start server", err)
	}
}
