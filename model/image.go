package model

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Owner     User
	Public    bool   `gorm:"default:false"`
	Favorites []User `gorm:"many2many:favorites;"`
}

func (i *Image) IsPublic() bool {
	return i.Public
}
