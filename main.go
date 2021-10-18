package main

import (
	"fmt"

	"github.com/gebhartn/impress/aws"
	"github.com/gebhartn/impress/db"
	"github.com/gebhartn/impress/handler"
	"github.com/gebhartn/impress/router"
	"github.com/gebhartn/impress/store"
)

func main() {
	a := aws.New()
	d := db.New()

	db.AutoMigrate(d)

	s3 := store.NewS3Store(a)
	user := store.NewUserStore(d)
	image := store.NewImageStore(d)

	r := router.New()
	h := handler.New(user, s3, image)

	h.Register(r)

	if err := r.Listen(":8080"); err != nil {
		fmt.Printf("%v", err)
	}
}
