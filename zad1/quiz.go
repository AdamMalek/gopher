package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := flag.String("input", "problems.csv", "input file containing problems with answers")
	timeLimit := flag.Int("time", 30, "time limit")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Error reading file %s: %s", *fileName, err.Error())
	}
	vals, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Printf("Error parsing file: %s", err.Error())
	}
	questions := make([]QuizItem, len(vals))
	for i, p := range vals {
		questions[i] = QuizItem{question: p[0], answer: p[1]}
	}

	consoleReader := bufio.NewReader(os.Stdin)
	answerChannel := make(chan string)
	getAnswer := func(question QuizItem) {
		ans, _ := consoleReader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		answerChannel <- ans
	}

	timeChannel := make(chan bool)
	go func() {
		time.Sleep(time.Duration(*timeLimit) * time.Second)
		timeChannel <- true
	}()

	correctAnswers := 0
	for i, q := range questions {
		fmt.Printf("%d. %s:\n", i+1, q.question)
		go getAnswer(q)
		select {
		case <-timeChannel:
			fmt.Printf("\nTotal score: %d/%d", correctAnswers, len(questions))
			return
		case ans := <-answerChannel:
			if ans == q.answer {
				correctAnswers++
			}
		}
	}
	fmt.Printf("Total score: %d/%d", correctAnswers, len(questions))
	fmt.Scanln()
}

type QuizItem struct {
	question string
	answer   string
}
