package main

import (
	"fmt"

	"github.com/cde/go-example/config"
	appError "github.com/cde/go-example/src/error"
	"github.com/cde/go-example/src/factory"
	"github.com/cde/go-example/src/handler"
	"github.com/cde/go-example/src/middleware"
	userFactory "github.com/cde/go-example/src/modules/user/factory"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: appError.CustomErrHandler,
	})

	// Middlewares
	app.Use(middleware.RequestId, middleware.LoggerContext, middleware.RequestLog)

	// load config
	cfg := config.Get()

	// Resolve dependencies
	validate = validator.New(validator.WithRequiredStructEnabled())
	db := factory.MakeGormDBConnection(cfg)
	userRepository := userFactory.ResolveUserRepository(db)
	userUseCase := userFactory.ResolveUserUseCase(userRepository)

	// register handler
	handler.NewUserHandler(app, validate, userUseCase)

	fmt.Println("Fiber app is running...")
	err := app.Listen(fmt.Sprintf(":%d", cfg.AppPort))
	if err != nil {
		panic(err)
	}
}
