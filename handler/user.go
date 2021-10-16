package handler

import (
	"net/http"

	"github.com/gebhartn/impress/model"
	"github.com/gebhartn/impress/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	if err := h.user.Create(&u); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newUserResponse(&u))
}
