package main

import (
	"log"

	"github.com/KentaKudo/go-todo/pkg/httpserver"
	"github.com/KentaKudo/go-todo/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Version string
)

func main() {
	db, err := mysql.OpenFromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatal(httpserver.New(db).Routes().Run(":8080"))
}
