package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	// Dummy auth logic
	return c.Next()
}
