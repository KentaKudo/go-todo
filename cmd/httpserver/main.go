package main

import (
	"log"

	skel "github.com/KentaKudo/goapi-skel"
	"github.com/KentaKudo/goapi-skel/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := mysql.NewFromYaml("./dbconfig.yml")
	if err != nil {
		log.Fatalln(err)
	}

	ts := mysql.NewTodoService(client)
	log.Fatal(skel.New(ts).Routes().Run(":8080"))
}
