package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lhdv/gophercises/07_task/model"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del [task id]",
	Short: "Deletes a task from your list",
	Long:  "Remove completly a task from your list, no matter if it was done or not.",
	Run: func(cmd *cobra.Command, args []string) {

		ids := make([]int, len(args))
		for k, arg := range args {
			var err error
			ids[k], err = strconv.Atoi(arg)
			if err != nil {
				log.Fatalln("Invalid arg:", arg)
				return
			}
		}

		ss, err := model.NewStorageService(model.WithBoltDB(""),
			model.WithBucket(""),
			model.WithTask())
		if err != nil {
			log.Fatalln(err)
		}

		for _, id := range ids {
			err := ss.Task.Delete(id)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Task", id, "was deleted from your list")
		}

		ss.Close()
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
