package middleware

import (
	"github.com/Ege-Okyay/go-file-server/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: config.AllowedOrigin,
	})
}
