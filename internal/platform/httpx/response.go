package httpx

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type errorBody struct {
	Error string `json:"error"`
}

// Error membalas dengan format error JSON yang konsisten.
func Error(c fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(errorBody{Error: msg})
}

// ErrorHandler dipasang di fiber.Config supaya semua error
// (termasuk panic yang ditangkap recover) berformat JSON.
func ErrorHandler(c fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	msg := "terjadi kesalahan"

	var fe *fiber.Error
	if errors.As(err, &fe) {
		status = fe.Code
		msg = fe.Message
	} else {
		slog.Error("error tidak tertangani", "error", err, "path", c.Path())
	}
	return Error(c, status, msg)
}
