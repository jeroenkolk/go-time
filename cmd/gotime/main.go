package main

import (
	"go-time/internal/db"
	"go-time/internal/http_handler"
)

func main() {
	db.SetupDB()
	http_handler.SetupWebServer()
}
