package main

import (
	"errors"
	"fmt"

	"github.com/cde/go-example/config"
	appError "github.com/cde/go-example/src/error"
	"github.com/cde/go-example/src/factory"
	"github.com/cde/go-example/src/handler"
	userFactory "github.com/cde/go-example/src/modules/user/factory"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err == nil {
				return nil
			}

			// Status code defaults to 500
			codeErrMessage := appError.CodeErrGeneral.GetCodeErrMessage()
			codeErrMessage.ErrDetail = err.Error()

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				codeErrMessage.StatusCode = e.Code
				return codeErrMessage.ToJson(ctx)
			}

			var codeErr appError.CodeErr
			if errors.As(err, &codeErr) {
				return codeErr.GetCodeErrMessage().ToJson(ctx)
			}

			var errMessage appError.CodeErrMessage
			if errors.As(err, &errMessage) {
				return errMessage.ToJson(ctx)
			}

			// another error
			return codeErrMessage.ToJson(ctx)
		},
	})

	// Logging Request ID
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	// load config
	config.LoadConfig()

	// Resolve dependencies
	validate = validator.New(validator.WithRequiredStructEnabled())
	db := factory.MakeGormDBConnection()
	userRepository := userFactory.ResolveUserRepository(db)
	userUseCase := userFactory.ResolveUserUseCase(userRepository)

	// register handler
	handler.NewUserHandler(app, validate, userUseCase)

	fmt.Println("Fiber app is running...")
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
