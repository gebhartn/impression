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
	us := store.NewUserStore(d)

	r := router.NewRouter()
	h := handler.NewHandler(us, s3)

	h.Register(r)

	if err := r.Listen(":8080"); err != nil {
		fmt.Printf("%v", err)
	}
}
