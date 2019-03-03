package cmd

import (
	"fmt"
	"strconv"

	"github.com/KentaKudo/go-todo/pkg/httpclient"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the todo",
	Long:  `Delete the todo`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("todo id is requiredtodo ")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		svc := httpclient.NewTodoService()
		if err := svc.Delete(id); err != nil {
			return err
		}

		return nil
	},
}
