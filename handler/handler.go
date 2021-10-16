package handler

type Handler struct {
	validator *Validator
}

func NewHandler() *Handler {
	v := NewValidator()
	return &Handler{
		validator: v,
	}
}
