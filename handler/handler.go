package handler

import "github.com/gebhartn/impress/user"

type Handler struct {
	user      user.Store
	validator *Validator
}

func NewHandler(us user.Store) *Handler {
	v := NewValidator()

	return &Handler{
		user:      us,
		validator: v,
	}
}
