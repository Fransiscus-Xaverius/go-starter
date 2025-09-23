package handler

import (
	"strconv"

	"github.com/cde/go-example/core/context"
	appError "github.com/cde/go-example/core/error"
	"github.com/cde/go-example/core/presentation"
	"github.com/cde/go-example/src/modules/user/dto"
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
	handler.registerEndpoints(app)
}

func (u *UserHandler) registerEndpoints(app *fiber.App) {
	app.Post("/users", u.Create)
	app.Get("/users/:id", u.GetByID)
	app.Get("/users", u.List)
}

func (u *UserHandler) Create(c *fiber.Ctx) error {
	var (
		logger, ctx = context.LoggerWithContext(c)
		req         dto.UserRequest
	)

	logger.Info("UserHandler.Create")
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := u.validate.Struct(req)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newUser, err := u.userUseCase.CreateUser(ctx, &req)
	if err != nil {
		return err
	}

	return presentation.NewResponseBuilder[dto.UserResponse]().
		SetData(dto.UserResponse{
			ID:    newUser.ID,
			Name:  newUser.Name,
			Email: newUser.Email,
		}).
		Json(c)
}

func (u *UserHandler) GetByID(c *fiber.Ctx) error {
	var (
		logger, ctx = context.LoggerWithContext(c)
	)

	logger.Info("UserHandler.GetByID")
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return appError.CodeErrValidation.WithErrorDetail(err)
	}

	user, err := u.userUseCase.GetUser(ctx, int32(id))
	if err != nil {
		return err
	}

	return presentation.NewResponseBuilder[*dto.UserResponse]().
		SetData(user).
		Json(c)
}

func (u *UserHandler) List(c *fiber.Ctx) error {
	var (
		logger, ctx = context.LoggerWithContext(c)
	)

	logger.Info("UserHandler.List")
	users, err := u.userUseCase.ListUsers(ctx, 10, 0)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return presentation.NewResponseBuilder[[]dto.UserResponse]().
		SetData(users).
		Json(c)
}
