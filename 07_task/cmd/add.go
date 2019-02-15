package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a task to your list",
	Long:  "Add a new task on top of your list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command", strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
