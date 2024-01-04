package apperror

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return NewAppError(404, message)
}

func NewBadRequestError(message string) *AppError {
	return NewAppError(400, message)
}

func NewInternalError(message string) *AppError {
	return NewAppError(500, message)
}

func NewUnauthorizedError(message string) *AppError {
	return NewAppError(401, message)
}

func NewForbiddenError(message string) *AppError {
	return NewAppError(403, message)
}

// ErrorHandler is a middleware that converts AppError to fiber.Error.
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Convert error to fiber.Error
	var fiberErr *fiber.Error
	if e, ok := err.(*AppError); ok {
		fiberErr = fiber.NewError(e.Code, e.Message)
	} else {
		// Default to 500 internal server error
		fiberErr = fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Send the error response
	return ctx.Status(fiberErr.Code).SendString(fiberErr.Message)
}
