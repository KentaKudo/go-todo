package main

import (
	"log"

	skel "github.com/KentaKudo/goapi-skel"
	"github.com/KentaKudo/goapi-skel/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client := mysql.NewClient(mysql.Config{})
	ts := mysql.NewTodoService(client)
	log.Fatal(skel.New(ts).Routes().Run(":8080"))
}
