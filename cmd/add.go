package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// add.goCmd represents the add.go command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your lists of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added %s to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
