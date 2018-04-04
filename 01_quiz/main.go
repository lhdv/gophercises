package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var quizFile string
	var timeOut int

	flag.StringVar(&quizFile, "f", "problems.csv", "Path to questions file(CSV format)")
	flag.IntVar(&timeOut, "t", 30, "How many seconds you have to realize an answer")
	flag.Parse()

	f := openFile(quizFile)
	qq := parseQuestions(f)

	fmt.Println("")
	fmt.Println("=================================")
	fmt.Println("===== Gopher Exercises Quiz =====")
	fmt.Println("=================================")
	fmt.Println("Are you ready?!?! (press return to continue)")
	fmt.Scanln()

	play(qq, timeOut)
}

func openFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal("[ERR] Error on openning file:", name, ". Err:", err)
	}

	return f
}

func parseQuestions(f *os.File) [][]string {
	csvFile := csv.NewReader(f)
	qq, err := csvFile.ReadAll()
	if err != nil {
		log.Fatal("[ERR] Error on parsing CSV file. Err:", err)
	}

	for _, q := range qq {
		q[0] = strings.TrimSpace(q[0])

		q[1] = strings.TrimSpace(q[1])
		q[1] = strings.ToUpper(q[1])
	}

	return qq
}

func play(quiz [][]string, timeOut int) {
	score := 0
	maxScore := len(quiz)

	gameOverFunc := func() {
		gameOver(score, maxScore)
		os.Exit(0)
	}

	time.AfterFunc(time.Duration(timeOut)*time.Second, gameOverFunc)

	for n, q := range quiz {
		var answer string
		fmt.Printf("Question %02d: %s = ", n+1, q[0])
		fmt.Scanln(&answer)

		answer = strings.TrimSpace(answer)
		answer = strings.ToUpper(answer)

		if answer == q[1] {
			score++
		}
	}

	gameOver(score, maxScore)
}

func gameOver(score, max int) {
	fmt.Println("\nYou scored", score, "of", max)
}
