package main

import (
	"log"

	"github.com/KentaKudo/goapi-skel/pkg/httpserver"
	"github.com/KentaKudo/goapi-skel/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := mysql.OpenFromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatal(httpserver.New(db).Routes().Run(":8080"))
}
