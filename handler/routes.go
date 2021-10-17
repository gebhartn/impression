package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")

	guests := v1.Group("/users")
	guests.Post("", h.SignUp)
	guests.Post("/login", h.Login)
}
