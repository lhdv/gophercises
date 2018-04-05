package main

import (
	"fmt"

	cyoa "github.com/lhdv/gophercises/03_cyoa"
)

func main() {
	fmt.Println("")
	fmt.Println("***********************************")
	fmt.Println("* Create Your Own Adventure - 1.0 *")
	fmt.Println("***********************************")
	fmt.Println("")

	cyoa.LoadAdventure("gopher.json")
}
