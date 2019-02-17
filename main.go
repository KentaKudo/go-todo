package main

import (
	"log"

	"github.com/KentaKudo/goapi-skel/pkg/httpserver"
	"github.com/KentaKudo/goapi-skel/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// TOOD
// - Swagger
// - tests

func main() {
	client, err := mysql.NewFromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	ts := mysql.NewTodoService(client)
	log.Fatal(httpserver.New(ts).Routes().Run(":8080"))
}
