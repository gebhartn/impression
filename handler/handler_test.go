package handler

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gebhartn/impress/aws"
	"github.com/gebhartn/impress/db"
	"github.com/gebhartn/impress/model"
	"github.com/gebhartn/impress/router"
	"github.com/gebhartn/impress/s3"
	"github.com/gebhartn/impress/store"
	"github.com/gebhartn/impress/user"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

var (
	d   *gorm.DB
	a   *session.Session
	us  user.Store
	s3s s3.Store
	h   *Handler
	e   *fiber.App
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	d = db.TestDB()
	a = aws.TestSession()

	db.AutoMigrate(d)

	us = store.NewUserStore(d)
	s3s = store.NewS3Store(a)

	h = New(us, s3s)
	e = router.New()

	loadFixures()
}

func tearDown() {
	_ = d.Close()
	if err := db.Drop(); err != nil {
		log.Fatal(err)
	}
}

func responseMap(bs []byte, key string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(bs, &m)
	return m[key].(map[string]interface{})
}

func loadFixures() error {
	u := model.User{
		Username: "testuser",
	}
	u.Password, _ = u.HashPassword("testpass")

	if err := us.Create(&u); err != nil {
		return err
	}

	return nil
}

func authHeader(token string) string {
	return "Token " + token
}
