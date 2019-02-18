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

		db, err := model.OpenDatabase("")
		if err != nil {
			log.Fatalln(err)
		}

		t, err := db.Add(strings.Join(args, " "), time.Now())
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("[CMD]", t)

		tt, err := db.Get(1)
		log.Println("[CMD GET]", tt)

		db.Close()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
