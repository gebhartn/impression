package db

import (
	"fmt"
	"os"

	"github.com/gebhartn/impress/model"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./database/impress.db")

	if err != nil {
		fmt.Println("[ERROR] database error", err)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../database/test.db")

	if err != nil {
		fmt.Println("[ERROR] NICK! database error", err)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	return db
}

func Drop() error {
	if err := os.Remove("./../database/test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Image{},
	)
}
