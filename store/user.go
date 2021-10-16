package store

import (
	"github.com/gebhartn/impress/model"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetById(id uint) (*model.User, error) {
	var u model.User
	if err := s.db.First(&u, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (s *UserStore) GetByUsername(name string) (*model.User, error) {
	var u model.User
	if err := s.db.Where(&model.User{Username: name}).First(&u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (s *UserStore) Create(u *model.User) (err error) {
	return s.db.Create(u).Error
}

func (s *UserStore) Update(u *model.User) (err error) {
	return s.db.Model(u).Update(u).Error
}
