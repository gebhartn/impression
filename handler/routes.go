package handler

import (
	"github.com/gebhartn/impress/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")

	guests := v1.Group("/users")
	guests.Post("", h.SignUp)
	guests.Post("/login", h.Login)

	auth := jwtware.New(jwtware.Config{
		SigningKey: utils.JWTSecret,
		AuthScheme: "Token",
	})

	user := v1.Group("/user", auth)
	user.Get("", h.CurrentUser)

	images := v1.Group("/images", auth)
	images.Post("/upload", h.UploadFile)
}
