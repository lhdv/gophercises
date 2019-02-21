package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/lhdv/gophercises/07_task/model"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Long:  "Show all your TODO list order by the last task added to the first",
	Run: func(cmd *cobra.Command, args []string) {

		var noTime time.Time

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

		fmt.Printf("\n======= TODO TASKS  =======\n\n")

		// List Todo tasks
		for i, t := range tasks {
			if t.CompletedAt == noTime {
				fmt.Printf("%02d. %s.\n", i+1, t.Desc)
			}
		}

		fmt.Printf("\n===== COMPLETED TASKS =====\n\n")

		// List Completed tasks
		completedIdx := 1
		for _, t := range tasks {
			if t.CompletedAt != noTime {
				fmt.Printf("%02d. %s. On %s\n", completedIdx, t.Desc, t.CompletedAt.Format("2006-01-02 15:04:05"))
				completedIdx++
			}
		}

		fmt.Println()

		ss.Close()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
