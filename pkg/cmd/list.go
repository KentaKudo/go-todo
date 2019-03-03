package cmd

import (
	"fmt"

	"github.com/KentaKudo/go-todo/pkg/httpclient"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the todo list",
	Long:  `Get the todo list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := httpclient.NewTodoService()
		todos, err := svc.List()
		if err != nil {
			return err
		}

		fmt.Printf("%v\n", todos)
		return nil
	},
}
