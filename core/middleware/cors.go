package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Cors = func(corsAllowOrigins string, allowHeaders string) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: corsAllowOrigins,
		AllowHeaders: allowHeaders,
	})
}
