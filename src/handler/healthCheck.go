package handler

import (
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
	app *fiber.App
}

func NewHealthCheckHandler(app *fiber.App) *HealthCheckHandler {
	h := &HealthCheckHandler{app: app}
	h.registerEndpoints(app)
	return h
}

func (h *HealthCheckHandler) registerEndpoints(app *fiber.App) {
	app.Get("/", h.HealthCheck)
	app.Get("/health", h.HealthCheck)
}

func (h *HealthCheckHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("OK")
}
