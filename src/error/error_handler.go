package error

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var CustomErrHandler = func(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	// Status code defaults to 500
	codeErrMessage := CodeErrGeneral.GetCodeErrMessage()
	codeErrMessage.ErrDetail = err.Error()

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		codeErrMessage.StatusCode = e.Code
		return codeErrMessage.ToJson(ctx)
	}

	var codeErr CodeErr
	if errors.As(err, &codeErr) {
		return codeErr.GetCodeErrMessage().ToJson(ctx)
	}

	var errMessage CodeErrMessage
	if errors.As(err, &errMessage) {
		return errMessage.ToJson(ctx)
	}

	// another error
	return codeErrMessage.ToJson(ctx)
}
