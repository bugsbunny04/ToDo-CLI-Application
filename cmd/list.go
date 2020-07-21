package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// list.goCmd represents the list.go command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
