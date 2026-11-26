package attendance

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v3"

	"mypulse/internal/platform/httpx"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Routes(r fiber.Router) {
	r.Post("/check-in", h.checkIn)
}

type checkInRequest struct {
	EmployeeID int64   `json:"employee_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func (h *Handler) checkIn(c fiber.Ctx) error {
	var req checkInRequest
	if err := c.Bind().Body(&req); err != nil {
		return httpx.Error(c, fiber.StatusBadRequest, "body tidak valid")
	}
	if req.EmployeeID <= 0 {
		return httpx.Error(c, fiber.StatusBadRequest, "employee_id wajib diisi")
	}

	a, err := h.svc.CheckIn(c.Context(), req.EmployeeID, req.Latitude, req.Longitude)
	if err != nil {
		if errors.Is(err, ErrInvalidLocation) {
			return httpx.Error(c, fiber.StatusUnprocessableEntity, err.Error())
		}
		slog.Error("check-in gagal", "error", err)
		return httpx.Error(c, fiber.StatusInternalServerError, "terjadi kesalahan")
	}
	return c.Status(fiber.StatusCreated).JSON(a)
}
