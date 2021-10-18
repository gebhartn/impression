package handler

import (
	"github.com/gebhartn/impress/image"
	"github.com/gebhartn/impress/s3"
	"github.com/gebhartn/impress/user"
)

type Handler struct {
	user      user.Store
	s3        s3.Store
	img       image.Store
	validator *Validator
}

func New(us user.Store, s3 s3.Store, img image.Store) *Handler {
	v := NewValidator()

	return &Handler{
		user:      us,
		s3:        s3,
		img:       img,
		validator: v,
	}
}
