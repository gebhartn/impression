package handler

import (
	"fmt"
	"net/http"

	"github.com/gebhartn/impress/model"
	"github.com/gebhartn/impress/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

func (h *Handler) Login(c *fiber.Ctx) error {
	req := &userLoginRequest{}

	if err := req.bind(c, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	u, err := h.user.GetByUsername(req.User.Username)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	if u == nil {
		return c.Status(http.StatusForbidden).JSON(utils.AccessForbidden())
	}

	if !u.CheckPassword(req.User.Password) {
		fmt.Printf("wrong password %v", err)
		return c.Status(http.StatusForbidden).JSON(utils.AccessForbidden())
	}

	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}

func (h *Handler) CurrentUser(c *fiber.Ctx) error {
	u, err := h.user.GetById(userIdFromToken(c))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	if u == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}

	return c.Status(http.StatusOK).JSON(newUserResponse(u))
}

func userIdFromToken(c *fiber.Ctx) uint {
	var user *jwt.Token
	l := c.Locals("user")
	if l == nil {
		return 0
	}
	user = l.(*jwt.Token)
	id := uint(((user.Claims.(jwt.MapClaims)["id"]).(float64)))
	return id
}
