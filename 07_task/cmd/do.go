package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Completes a tasks",
	Long:  "Mark a task as completed and remove it from your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Do command")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
