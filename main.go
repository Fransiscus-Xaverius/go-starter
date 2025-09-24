package main

import (
	"fmt"

	"github.com/cde/go-example/config"
	appError "github.com/cde/go-example/core/error"
	"github.com/cde/go-example/core/factory"
	coreHandler "github.com/cde/go-example/core/handler"
	"github.com/cde/go-example/core/middleware"
	"github.com/cde/go-example/src/handler"
	userFactory "github.com/cde/go-example/src/modules/user/factory"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	app := fiber.New(fiber.Config{
		// Override default error coreHandler
		ErrorHandler: appError.CustomErrHandler,
	})

	// load config
	cfg := config.Get()

	// Middlewares
	app.Use(
		middleware.Cors(cfg.CorsAllowOrigins, cfg.AllowHeaders),
		middleware.RequestId,
		middleware.LoggerContext(cfg.AppName, cfg.AppVersion),
		middleware.RequestLog,
	)

	// Resolve dependencies
	validate = validator.New(validator.WithRequiredStructEnabled())
	db := factory.MakeGormDBConnection(cfg)
	userRepository := userFactory.ResolveUserRepository(db)
	userUseCase := userFactory.ResolveUserUseCase(userRepository)

	// register coreHandler
	coreHandler.NewHealthCheckHandler(app)

	// register appHandler
	handler.NewUserHandler(app, validate, userUseCase)

	fmt.Printf("%s app is running...\n", cfg.AppName)
	err := app.Listen(fmt.Sprintf(":%d", cfg.AppPort))
	if err != nil {
		panic(err)
	}
}
