package middleware

import (
	"github.com/cde/go-example/core/context"
	"github.com/gofiber/fiber/v2"
)

const (
	LoggerRequestId = "request_id"
)

func LoggerContext(appName string, appVersion string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		value := ctx.Value(context.RequestId{})
		c.Locals(
			context.Logger{},
			context.
				NewLogger().
				WithField(
					LoggerRequestId,
					value,
				).
				WithField(
					"app_name",
					appName,
				).
				WithField(
					"app_version",
					appVersion,
				))

		return c.Next()
	}
}
