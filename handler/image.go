package handler

import (
	"net/http"

	"github.com/gebhartn/impress/model"
	"github.com/gebhartn/impress/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UploadFile(c *fiber.Ctx) error {
	var i model.Image
	r := &imageUploadRequest{}

	if err := r.bind(c, &i, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	key, err := h.s3.UploadObject(r.Image.Owner, r.Image.File)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	i.Key = key

	return c.Status(http.StatusCreated).JSON(newImageUploadResponse(&i))
}
