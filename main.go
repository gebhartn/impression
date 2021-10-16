package main

import (
	"fmt"

	"github.com/gebhartn/impress/db"
	"github.com/gebhartn/impress/handler"
	"github.com/gebhartn/impress/router"
)

func main() {
	d := db.New()
	db.AutoMigrate(d)

	r := router.NewRouter()
	h := handler.NewHandler()

	h.Register(r)

	err := r.Listen(":8080")
	if err != nil {
		fmt.Printf("%v", err)
	}
}
