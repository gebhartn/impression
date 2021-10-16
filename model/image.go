package model

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Owner     User
	Public    bool
	Favorites []User `gorm:"many2many:favorites;"`
}

func (i *Image) IsPublic() bool {
	return i.Public
}
