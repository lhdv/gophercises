package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Long:  "Show all yout TODO list order by the last task added to the first",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List command")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
