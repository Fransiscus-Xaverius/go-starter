package middleware

import (
	"github.com/cde/go-example/core/context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

var RequestId = requestid.New(requestid.Config{
	Header:     fiber.HeaderXRequestID,
	Generator:  utils.UUIDv4,
	ContextKey: context.RequestId{},
})
