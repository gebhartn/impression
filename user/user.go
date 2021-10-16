package user

import "github.com/gebhartn/impress/model"

type Store interface {
	GetById(uint) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
}
