package handler

import (
	"strconv"

	appError "github.com/cde/go-example/src/error"
	userDTO "github.com/cde/go-example/src/modules/user/dto"
	"github.com/cde/go-example/src/modules/user/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase usecase.UserUseCaseInterface
	validate    *validator.Validate
}

func NewUserHandler(app *fiber.App, validate *validator.Validate, userUseCase usecase.UserUseCaseInterface) {
	handler := &UserHandler{userUseCase: userUseCase, validate: validate}
	registerEndpoints(app, handler)
}

func registerEndpoints(app *fiber.App, handler *UserHandler) {
	app.Post("/users", handler.Create)
	app.Get("/users/:id", handler.GetByID)
	app.Get("/users", handler.List)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req userDTO.UserRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := h.validate.Struct(req)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newUser, err := h.userUseCase.CreateUser(ctx, &req)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(userDTO.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	})
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return appError.CodeErrValidation.WithErrorDetail(err)
	}

	user, err := h.userUseCase.GetUser(c.Context(), int32(id))
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.userUseCase.ListUsers(c.Context(), 10, 0)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(users)
}
