package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")

	hello := v1.Group("/:name")
	hello.Get("", h.SayHello)
}

func (h *Handler) SayHello(c *fiber.Ctx) error {
	name := c.Params("name", "world")
	res := map[string]interface{}{"res": fmt.Sprintf("Hello, %s!", name)}
	return c.Status(http.StatusOK).JSON(res)
}
