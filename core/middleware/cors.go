package middleware

import (
	"github.com/cde/go-example/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Cors = func(cfg *config.Config) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: cfg.CorsAllowOrigins,
		AllowHeaders: cfg.AllowHeaders,
	})
}
