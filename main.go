package main

import (
	"fmt"

	"github.com/Ege-Okyay/go-file-server/config"
	"github.com/Ege-Okyay/go-file-server/handlers"
	"github.com/Ege-Okyay/go-file-server/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadConfig()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(middleware.SetupCORS())

	app.Static("/uploads", config.UploadPath)

	app.Post("/upload", handlers.UploadFile)

	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
