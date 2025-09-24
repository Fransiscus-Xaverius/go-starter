package error

import (
	"errors"

	"github.com/cde/go-example/core/presentation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var CustomErrHandler = func(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	// Status code defaults to 500
	codeErrMessage := CodeErrGeneral.GetCodeErr()
	codeErrMessage.Detail = err.Error()

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		codeErrMessage.StatusCode = e.Code
		return codeErrMessage.ToJson(ctx)
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var errList []string
		for _, err := range validationErrors {
			errList = append(errList, err.Error())
		}
		codeErr := CodeErrValidation.GetCodeErr()
		return presentation.Response[[]string]().
			SetStatus(false).
			SetErrorCode(codeErr.Code).
			SetMessage(codeErr.Message).
			SetStatusCode(fiber.StatusBadRequest).
			WithData(errList).
			Json(ctx)
	}

	var codeErr CodeErrEnum
	if errors.As(err, &codeErr) {
		return codeErr.GetCodeErr().ToJson(ctx)
	}

	var errMessage CodeErr
	if errors.As(err, &errMessage) {
		return errMessage.ToJson(ctx)
	}

	// another error
	return codeErrMessage.ToJson(ctx)
}
