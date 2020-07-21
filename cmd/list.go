package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// list.goCmd represents the list.go command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("no tasks ! Why not take a vacation! ")
			return
		}
		fmt.Println("You have the following tasks: \n")
		for i, task := range tasks {
			fmt.Printf("%d. %s... Key=%d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
