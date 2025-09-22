package error

import (
	"github.com/cde/go-example/config"
	"github.com/gofiber/fiber/v2"
)

type (
	CodeErr int

	CodeErrMessage struct {
		ErrMessage string `json:"error_message"`
		ErrDetail  string `json:"error_detail,omitempty"`
		ErrCode    string `json:"error_code"`
		StatusCode int    `json:"status_code,omitempty"`
	}
)

const (
	CodeErrGeneral CodeErr = iota
	CodeErrValidation
	CodeErrUserNotFound
)

var (
	codeErrMap = map[CodeErr]CodeErrMessage{
		CodeErrGeneral:      {ErrMessage: "Something went wrong", ErrCode: "ERRDEMO500001", StatusCode: fiber.StatusInternalServerError},
		CodeErrValidation:   {ErrMessage: "Validation error", ErrCode: "ERRDEMO400002", StatusCode: fiber.StatusBadRequest},
		CodeErrUserNotFound: {ErrMessage: "User not found", ErrCode: "ERRDEMO404003", StatusCode: fiber.StatusNotFound},
	}
)

func (c CodeErr) Error() string {
	return c.GetCodeErrMessage().ErrMessage
}

func (c CodeErr) GetCodeErrMessage() CodeErrMessage {
	if code, ok := codeErrMap[c]; ok {
		return code
	}
	return codeErrMap[CodeErrGeneral]
}

func (c CodeErr) WithErrorDetail(err error) CodeErrMessage {
	errMessage := c.GetCodeErrMessage()
	errMessage.ErrDetail = err.Error()
	return errMessage
}

func (c CodeErrMessage) Error() string {
	return c.ErrMessage
}

func (c CodeErrMessage) ToJson(ctx *fiber.Ctx) error {
	statusCode := c.StatusCode
	c.StatusCode = 0
	if !config.Get().AppDebug {
		c.ErrDetail = ""
	}
	return ctx.Status(statusCode).JSON(c)
}
