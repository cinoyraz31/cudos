package exceptions

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandlerResp struct {
	Message string      `json:"message"`
	Err     interface{} `json:"errors"`
}

type ErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrorHandlerInternalServerError(ctx *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			ErrorHandlerCustom(ctx, fiber.StatusInternalServerError, fmt.Sprintf("%v", r))
		}
	}()
	return ctx.Next()
}

func ErrorHandlerUnprocessableEntity(ctx *fiber.Ctx, err error) error {
	validationErrors := err.(validator.ValidationErrors)
	errorArr := make([]ErrorField, len(validationErrors))

	for i, err := range validationErrors {
		errorArr[i] = ErrorField{
			Field:   err.Field(),
			Message: err.Tag(),
		}
	}

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(ErrorHandlerResp{
		Message: err.Error(),
		Err:     errorArr,
	})
}

func ErrorHandlerCustom(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(ErrorHandlerResp{
		Message: message,
	})
}

func ErrorHandlerBadRequest(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(ErrorHandlerResp{
		Message: message,
	})
}
