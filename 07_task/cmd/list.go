package cmd

import (
	"fmt"
	"log"

	"github.com/lhdv/gophercises/07_task/model"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Long:  "Show all yout TODO list order by the last task added to the first",
	Run: func(cmd *cobra.Command, args []string) {

		ss, err := model.NewStorageService(model.WithBoltDB(""),
			model.WithBucket(""),
			model.WithTask())
		if err != nil {
			log.Fatalln(err)
		}

		ts, err := ss.Task.List()
		if err != nil {
			log.Fatalln(err)
		}

		for i, t := range ts {
			fmt.Printf("%d. %+v\n", i+1, t)
		}

		ss.Close()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
