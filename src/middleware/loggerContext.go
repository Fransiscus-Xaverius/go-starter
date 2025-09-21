package middleware

import (
	"github.com/cde/go-example/src/context"
	"github.com/gofiber/fiber/v2"
)

func LoggerContext(c *fiber.Ctx) error {
	ctx := c.Context()
	builder := context.NewContextBuilder(ctx).
		SetLogger(
			context.
				NewLogger().
				WithField(
					"request_id",
					ctx.Value(context.RequestId{}),
				),
		)
	c.Locals(builder.Context())
	return c.Next()
}
