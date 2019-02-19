package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/lhdv/gophercises/07_task/model"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Completes tasks",
	Long:  "Mark one or more tasks as completed",
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

		tasksIDs := make([]int, 0, len(tasks))
		for _, id := range ids {
			if id <= 0 && id > len(tasks) {
				log.Println("Invalid id:", id)
				continue
			}

			task := tasks[id-1]
			if task.ID != 0 {
				tasksIDs = append(tasksIDs, task.ID)
			}

		}

		completedTasks, err := ss.Task.Complete(tasksIDs, time.Now())
		if err != nil {
			log.Fatalln(err)
		}

		for _, t := range completedTasks {
			fmt.Printf("Task \"%s\" set as completed on %q\n", t.Desc, time.Now().Format("02/01/2006 15:04:05"))
		}

		ss.Close()
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
