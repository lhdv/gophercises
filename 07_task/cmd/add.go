package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/lhdv/gophercises/07_task/model"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a task to your list",
	Long:  "Add a new task on top of your list",
	Run: func(cmd *cobra.Command, args []string) {

		ss, err := model.NewStorageService(model.WithBoltDB(""),
			model.WithBucket(""),
			model.WithTask())
		if err != nil {
			log.Fatalln(err)
		}

		_, err = ss.Task.Add(strings.Join(args, " "), time.Now())
		if err != nil {
			log.Fatalln(err)
		}

		ss.Close()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
