package cmd

import (
	"fmt"
	"strconv"

	"github.com/KentaKudo/go-todo/pkg/httpclient"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single todo",
	Long:  "Get a single todo",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("todo id is required")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		svc := httpclient.NewTodoService()
		todo, err := svc.Get(id)
		if err != nil {
			return err
		}

		fmt.Printf("%v\n", todo)
		return nil
	},
}
