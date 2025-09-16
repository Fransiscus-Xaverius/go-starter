package delivery

import (
    "github.com/gofiber/fiber/v2"
	userDTO "go-example/internal/modules/user/dto"
	generalDTO "go-example/internal/general/dto"
	usecase "go-example/internal/modules/user/usecase"
	// middleware "go-example/internal/middleware"
)

type UserHandler struct {
    userUsecase usecase.UserUseCase
}

func NewUserHandler(app *fiber.App, uc usecase.UserUseCase) {
    handler := &UserHandler{userUsecase: uc}
    app.Post("/users", handler.Create)
    app.Get("/users/:id", handler.GetByID)
    app.Get("/users", handler.List)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
    var req struct {
        Name  string `json:"name"` 
        Email string `json:"email"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    newUser, err := h.userUsecase.CreateUser(req.Name, req.Email)
    if err != nil {
        return c.Status(500).JSON(generalDTO.ErrorResponse{
            Message: err.Error(),
            Details: "Error Details here",
            ErrorCode: "CREATE-1",
        })
    }
    return c.Status(201).JSON(userDTO.UserResponse{
        ID:    newUser.ID,
        Name:  newUser.Name,
        Email: newUser.Email,
    })
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
    id := c.Params("id")
    user, err := h.userUsecase.GetUser(id)
    if err != nil {
        return c.Status(404).SendString(err.Error())
    }
    return c.JSON(user)
}

func (h *UserHandler) List(c *fiber.Ctx) error {
    users, err := h.userUsecase.ListUsers()
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(users)
}