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

		tasks, err := ss.Task.List()
		if err != nil {
			log.Fatalln(err)
		}

		for _, id := range ids {
			if id <= 0 && id > len(tasks) {
				log.Println("Invalid id:", id)
				continue
			}

			task := tasks[id-1]
			if task.ID != 0 {
				err := ss.Task.Delete(task.ID)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Printf("Task %02d - %s was deleted from your list\n", id, task.Desc)
			}

		}

		ss.Close()
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
