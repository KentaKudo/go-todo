package main

import (
	"log"

	"github.com/KentaKudo/goapi-skel/httpserver"
	"github.com/KentaKudo/goapi-skel/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// TOOD
// - debug MySQL CRUD interface
// - Dockerise
// - docker-compose
// - Swagger
// - tests

func main() {
	client, err := mysql.NewFromYaml("./dbconfig.yml")
	if err != nil {
		log.Fatalln(err)
	}

	ts := mysql.NewTodoService(client)
	log.Fatal(httpserver.New(ts).Routes().Run(":8080"))
}
