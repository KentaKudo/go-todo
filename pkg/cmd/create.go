package cmd

import (
	"fmt"

	todo "github.com/KentaKudo/go-todo"
	"github.com/KentaKudo/go-todo/pkg/httpclient"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new todo",
	Long:  `Create a new todo`,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := todo.Todo{Title: "new todo"}
		svc := httpclient.NewTodoService()
		if err := svc.Create(&t); err != nil {
			return err
		}

		fmt.Printf("%v\n", t)
		return nil
	},
}
