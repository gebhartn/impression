package handler

import (
	"fmt"
	"net/http"
	"strings"

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

	i.Key = extractImageName(key)
	i.OwnerID = r.Image.Owner

	if err = h.img.Create(&i); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newImageUploadResponse(&i))
}

func (h *Handler) GetUploads(c *fiber.Ctx) error {
	is, err := h.s3.ListObjectsById(userIdFromToken(c))
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if len(is.Contents) == 1 {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}

	fmt.Printf("\n%v\n", is)

	return c.Status(http.StatusOK).JSON((newUserImagesReponse(is)))
}

func extractImageName(url string) string {
	l := strings.Split(url, "/")
	return l[len(l)-1]
}
