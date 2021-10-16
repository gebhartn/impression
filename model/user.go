package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique_index;not null"`
	Password  string `gorm:"not null"`
	Uploads   []Image
	Favorites []Image `gorm:"many2many:favorites;"`
}

func (u *User) HashPassword(pw string) (string, error) {
	if len(pw) == 0 {
		return "", errors.New("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	return string(hash), err
}

func (u *User) CheckPassword(pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
	return err == nil
}
