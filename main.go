package main

import (
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
		ErrorHandler: appError.CustomErrHandler,
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
