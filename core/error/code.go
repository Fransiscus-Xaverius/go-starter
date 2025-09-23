package error

import (
	"github.com/cde/go-example/config"
	"github.com/cde/go-example/core/presentation/dto"
	"github.com/gofiber/fiber/v2"
)

type (
	CodeErrEnum int

	CodeErr struct {
		Message    string
		Detail     string
		Code       string
		StatusCode int
	}
)

const (
	CodeErrGeneral CodeErrEnum = iota
	CodeErrValidation
	CodeErrUserNotFound
	CodeErrUnauthorizedEmptyAuthorizationHeader
	CodeErrUnauthorizedInvalidAuthorizationHeader
	CodeErrUnauthorizedNonBearerAuthorizationHeader
	CodeErrUnauthorizedEmptyBearerAuthorizationHeader
)

var (
	codeErrMap = map[CodeErrEnum]CodeErr{
		CodeErrGeneral:      {Message: "Something went wrong", Code: "ERRDEMO500001", StatusCode: fiber.StatusInternalServerError},
		CodeErrValidation:   {Message: "Validation error", Code: "ERRDEMO400002", StatusCode: fiber.StatusBadRequest},
		CodeErrUserNotFound: {Message: "User not found", Code: "ERRDEMO404003", StatusCode: fiber.StatusNotFound},
		CodeErrUnauthorizedEmptyAuthorizationHeader:       {Message: "Unauthorized", Code: "ERRDEMO401004", StatusCode: fiber.StatusUnauthorized},
		CodeErrUnauthorizedInvalidAuthorizationHeader:     {Message: "Unauthorized", Code: "ERRDEMO401005", StatusCode: fiber.StatusUnauthorized},
		CodeErrUnauthorizedNonBearerAuthorizationHeader:   {Message: "Unauthorized", Code: "ERRDEMO401006", StatusCode: fiber.StatusUnauthorized},
		CodeErrUnauthorizedEmptyBearerAuthorizationHeader: {Message: "Unauthorized", Code: "ERRDEMO401007", StatusCode: fiber.StatusUnauthorized},
	}
)

func (c CodeErrEnum) Error() string {
	return c.GetCodeErr().Message
}

func (c CodeErrEnum) GetCodeErr() CodeErr {
	if code, ok := codeErrMap[c]; ok {
		return code
	}
	return codeErrMap[CodeErrGeneral]
}

func (c CodeErrEnum) WithErrorDetail(err error) CodeErr {
	errMessage := c.GetCodeErr()
	errMessage.Detail = err.Error()
	return errMessage
}

func (c CodeErr) Error() string {
	return c.Message
}

func (c CodeErr) ToJson(ctx *fiber.Ctx) error {
	data := dto.Response[any]{
		Status:    false,
		Message:   c.Message,
		ErrDetail: c.Detail,
		ErrCode:   c.Code,
	}
	if !config.Get().AppDebug {
		data.ErrDetail = ""
	}
	return ctx.Status(c.StatusCode).JSON(data)
}
