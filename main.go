package main

import (
	"fmt"

	"github.com/cde/go-example/config"
	"github.com/cde/go-example/src/handler"
	userFactory "github.com/cde/go-example/src/modules/user/factory"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	app := fiber.New()
	config.LoadConfig()
	//db := factory.MakeGormDBConnection()

	validate = validator.New(validator.WithRequiredStructEnabled())

	// Resolve dependencies
	userUseCase := userFactory.ResolveUserUseCase()

	// register handler
	handler.NewUserHandler(app, validate, userUseCase)

	fmt.Println("Fiber app is running...")
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
