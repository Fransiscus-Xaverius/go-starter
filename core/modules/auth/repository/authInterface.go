package repository

import (
	"github.com/cde/go-example/core/modules/auth/dto"
	"github.com/gofiber/fiber/v2"
)

type AuthRepositoryInterface interface {
	Authorize(c *fiber.Ctx, token string) (dto.AuthorizeResponse, error)
}
