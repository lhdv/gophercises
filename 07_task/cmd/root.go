package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task manages your TODO list",
	Long:  "Task manages and stores your TODO list to keep you in control of your life",
}

// Execute parses and runs your CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
