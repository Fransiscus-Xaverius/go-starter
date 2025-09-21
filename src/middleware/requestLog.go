package middleware

import (
	"fmt"
	"time"

	"github.com/cde/go-example/src/context"
	"github.com/gofiber/fiber/v2"
)

func RequestLog(c *fiber.Ctx) error {
	var (
		logger = context.LoggerFromContext(c.Context())
	)

	start := time.Now()

	// Log Request
	logger.Info(fmt.Sprintf("Method: %s Path:%s Headers: %v", c.Method(), c.Path(), c.GetReqHeaders()))

	// Proceed to next middleware/handler
	err := c.Next()

	// Log Response
	statusCode := c.Response().StatusCode()
	responseMessage := fmt.Sprintf("Status: %d Payload: %s Duration: %v", statusCode, string(c.Response().Body()), time.Since(start))
	if statusCode >= 200 && statusCode < 400 {
		logger.Info(responseMessage)
	} else if statusCode >= 400 && statusCode < 500 {
		logger.Warning(responseMessage)
	} else {
		logger.Error(responseMessage)
	}

	return err
}
