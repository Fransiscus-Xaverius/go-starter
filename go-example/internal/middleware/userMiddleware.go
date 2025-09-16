package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
    // Dummy auth logic
    return c.Next()
}

func MiddlewareRequestContextDeadline(c *fiber.Ctx) error {
    // Dummy context deadline 
    return c.Next()
}

func MaximumRequestSize(c *fiber.Ctx) error {
    // Dummy maximum request size
    return c.Next()
}

func CORS(c *fiber.Ctx) error {
    // Dummy CORS
    return c.Next()
}