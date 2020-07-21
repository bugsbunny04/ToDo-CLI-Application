package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// do.goCmd represents the do.go command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// do.goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// do.goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
