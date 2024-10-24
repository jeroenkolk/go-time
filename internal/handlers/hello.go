package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, Go Time!")
}

type TimeResponse struct {
	Time string `json:"time"`
}

func GetTime(c *fiber.Ctx) error {
	tz := c.Query("timezone", "Europe/Amsterdam")

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid timezone")
	}

	currentTime := time.Now().In(loc).Format("02-01-2006 15:04:05")

	response := TimeResponse{
		Time: currentTime,
	}

	return c.JSON(response)
}
