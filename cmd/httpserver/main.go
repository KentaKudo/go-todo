package main

import (
	"log"

	"github.com/KentaKudo/goapi-skel/mock"

	skel "github.com/KentaKudo/goapi-skel"
)

func main() {
	ts := mock.NewTodoService(func() ([]skel.Todo, error) {
		return []skel.Todo{skel.Todo{Title: "Hello, world"}}, nil
	})
	log.Fatal(skel.New(ts).Routes().Run(":8080"))
}
