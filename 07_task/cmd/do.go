package cmd

import (
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

		db, err := model.OpenDatabase("")
		if err != nil {
			log.Fatalln(err)
		}

		t, err := db.Complete(ids, time.Now())
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(t)

		db.Close()
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
