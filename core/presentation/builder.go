package presentation

import (
	"github.com/cde/go-example/core/presentation/dto"
	"github.com/gofiber/fiber/v2"
)

type (
	ResponseBuilder[T any] struct {
		statusCode int
		response   dto.Response[T]
	}
)

func NewResponseBuilder[T any]() *ResponseBuilder[T] {
	return &ResponseBuilder[T]{
		statusCode: fiber.StatusOK,
		response: dto.Response[T]{
			Status: true,
		},
	}
}

func (b *ResponseBuilder[T]) SetStatus(status bool) *ResponseBuilder[T] {
	b.response.Status = status
	return b
}

func (b *ResponseBuilder[T]) SetMessage(msg string) *ResponseBuilder[T] {
	b.response.Message = msg
	return b
}

func (b *ResponseBuilder[T]) SetErrorDetail(detail string) *ResponseBuilder[T] {
	b.response.ErrDetail = detail
	return b
}

func (b *ResponseBuilder[T]) SetErrorCode(code string) *ResponseBuilder[T] {
	b.response.ErrCode = code
	return b
}

func (b *ResponseBuilder[T]) SetData(data T) *ResponseBuilder[T] {
	b.response.Data = data
	return b
}

func (b *ResponseBuilder[T]) SetStatusCode(statusCode int) *ResponseBuilder[T] {
	b.statusCode = statusCode
	return b
}

func (b *ResponseBuilder[T]) Build() dto.Response[T] {
	return b.response
}

func (b *ResponseBuilder[T]) Json(ctx *fiber.Ctx) error {
	return ctx.Status(b.statusCode).JSON(b.Build())
}
