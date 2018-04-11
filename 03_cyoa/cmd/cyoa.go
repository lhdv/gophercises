package main

import (
	"flag"
	"fmt"

	"github.com/lhdv/gophercises/03_cyoa/server"
)

func main() {

	var inputFile string
	var srvHost string
	var srvPort string

	flag.StringVar(&inputFile, "f", "gopher.json", "input file which contains the adventure to load in json format")
	flag.StringVar(&srvHost, "s", "localhost", "the name/ip of the host which will be the application server")
	flag.StringVar(&srvPort, "p", "8080", "the port of the host which will be the application server")
	flag.Parse()

	fmt.Println("")
	fmt.Println("***********************************")
	fmt.Println("* Create Your Own Adventure - 1.0 *")
	fmt.Println("***********************************")
	fmt.Println("")

	// cyoa.LoadAdventure(inputFile)
	server.Start(inputFile, srvHost, srvPort)
}
