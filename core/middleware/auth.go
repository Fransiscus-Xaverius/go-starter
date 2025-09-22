package middleware

import (
	"strings"

	appContext "github.com/cde/go-example/core/context"
	appError "github.com/cde/go-example/core/error"
	"github.com/cde/go-example/core/modules/auth/repository"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(authRepository repository.AuthRepositoryInterface) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authStr := c.Get(fiber.HeaderAuthorization)
		if authStr == "" {
			return appError.CodeErrUnauthorizedEmptyAuthorizationHeader
		}
		authParts := strings.Split(authStr, " ")
		if len(authParts) != 2 {
			return appError.CodeErrUnauthorizedInvalidAuthorizationHeader
		}
		if authParts[0] == "Bearer" {
			return appError.CodeErrUnauthorizedNonBearerAuthorizationHeader
		}
		tokenStr := authParts[1]
		if tokenStr == "" {
			return appError.CodeErrUnauthorizedNonBearerAuthorizationHeader
		}

		ctx := c.Context()
		authorizeResponse, err := authRepository.Authorize(ctx, tokenStr)
		if err != nil {
			return err
		}

		builder := appContext.NewContextBuilder(ctx).
			SetSession(
				authorizeResponse,
			)
		c.SetUserContext(builder.Context())

		return c.Next()
	}
}
