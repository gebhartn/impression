package store

import (
	"github.com/gebhartn/impress/model"
	"github.com/jinzhu/gorm"
)

type ImageStore struct {
	db *gorm.DB
}

func NewImageStore(db *gorm.DB) *ImageStore {
	return &ImageStore{
		db: db,
	}
}

func (s *ImageStore) Create(i *model.Image) (err error) {
	return s.db.Create(i).Error
}
