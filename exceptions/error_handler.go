package exceptions

import (
	"errors"
	"yordanluturyali/golang-auth-rest/model"
	"github.com/gofiber/fiber/v2"
)

var globalResponse = &model.GlobalResponse{
	Data: nil,
}

func HandleError(c *fiber.Ctx, err error) error {
	if errors.As(err, &FailedValidation{}) {
		return validationError(c, err.(FailedValidation))
	}

	if errors.As(err, &HttpInternalServerError{}) {
		return internalServerError(c, err.(HttpInternalServerError))
	}

	var e *fiber.Error
	if errors.As(err, &e) {
		globalResponse.Message = e.Message
		globalResponse.Errors = nil
		c.Status(e.Code).JSON(&globalResponse)
	}

	return baseError(c, err.(GlobalError))
}

func validationError(c *fiber.Ctx, err FailedValidation) error {
	globalResponse.Message = err.Error()
	globalResponse.Errors = err.Errors

	return c.Status(err.GetCode()).JSON(&globalResponse)
}

func baseError(c *fiber.Ctx, err GlobalError) error {
	globalResponse.Message = err.Error()
	globalResponse.Errors = nil

	return c.Status(err.GetCode()).JSON(&globalResponse)
}

func internalServerError(c *fiber.Ctx, err HttpInternalServerError) error {
	globalResponse.Message = err.Error()
	globalResponse.Errors = nil

	return c.Status(err.GetCode()).JSON(&globalResponse)
}