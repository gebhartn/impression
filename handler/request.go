package handler

import (
	"mime/multipart"
	"strconv"

	"github.com/gebhartn/impress/model"
	"github.com/gofiber/fiber/v2"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userRegisterRequest) bind(c *fiber.Ctx, u *model.User, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	u.Username = r.User.Username
	hash, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = hash

	return nil
}

type userLoginRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c *fiber.Ctx, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	return nil
}

type imageUploadRequest struct {
	Image struct {
		Owner uint
		File  *multipart.FileHeader
	}
}

func (r *imageUploadRequest) bind(c *fiber.Ctx, i *model.Image, v *Validator) error {
	fh, err := c.FormFile("image")
	if err != nil {
		return err
	}

	o, err := strconv.Atoi(c.Params("slug"))
	if err != nil {
		return err
	}

	r.Image.File = fh
	r.Image.Owner = uint(o)

	return nil
}
