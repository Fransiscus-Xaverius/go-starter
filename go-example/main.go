package main

import (
    "fmt"
    "go-example/internal/config"
    "go-example/internal/modules/user/delivery"
    "go-example/internal/modules/user/repository"
    "go-example/internal/modules/user/usecase"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    config.LoadConfig()
    
    userRepo := userRepository.NewMockRepo()
    userUC := user.NewUseCase(userRepo)
    delivery.NewUserHandler(app, userUC)
    
    fmt.Println("Fiber app is running...")
    app.Listen(":3000")
}
