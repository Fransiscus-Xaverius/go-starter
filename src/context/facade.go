package context

import (
	"context"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func LoggerWithContext(c *fiber.Ctx) (*log.Entry, context.Context) {
	builder := NewContextBuilder(c.Context())
	return builder.GetLogger(), builder.Context()
}

func LoggerFromContext(ctx context.Context) *log.Entry {
	return NewContextBuilder(ctx).GetLogger()
}
