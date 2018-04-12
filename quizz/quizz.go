// Package quizz This is the first exercise from goexercises.com
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Quiz data structure for quizzes
type Quiz struct {
	problemStatement string
	correctAnswer    string
}

// Attempt data structure to record attempts of a quizz
type Attempt struct {
	quiz       Quiz
	answer     string
	isTimedOut bool
}

const (
	// CORRECT contains unicode symbol for marking correct
	CORRECT = "✓"
	// INCORRECT contains unicode symbol for marking incorrect
	INCORRECT = "❌"
)

func readProblems(csvFile string) []Quiz {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal("Error in reading file", err)
	}
	csvReader := csv.NewReader(file)
	// make a slice of Quiz, this is the way to create dynamic size array
	quizzes := make([]Quiz, 0)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quizz := Quiz{record[0], record[1]}
		// append is returning a new slice
		quizzes = append(quizzes, quizz)
	}
	return quizzes
}

func printScore(quizzes []Quiz, attempts []Attempt) {
	fmt.Println("===Here is your attempt Summary and Score===")
	var score int
	for index, attempt := range attempts {
		if attempt.answer != attempt.quiz.correctAnswer {
			fmt.Printf("%d %s=%s %s . CorrectAnswer: %s\n",
				(index + 1),
				attempt.quiz.problemStatement,
				attempt.answer,
				INCORRECT,
				attempt.quiz.correctAnswer)
		} else {
			score++
			fmt.Printf("%d %s=%s %s\n",
				(index + 1),
				attempt.quiz.problemStatement,
				attempt.answer,
				CORRECT)
		}
	}
	fmt.Printf("Final Score: %d/%d\n\n", score, len(quizzes))
}

func askQuestion(quiz Quiz, index int, channel chan Attempt) {
	fmt.Printf("%d %s: ", (index + 1), quiz.problemStatement)
	var res string
	fmt.Scan(&res)
	channel <- Attempt{quiz: quiz, answer: res}
}

func main() {
	quizzes := readProblems("problems.csv")
	attempts := make([]Attempt, 0)
	fmt.Println("Please answer to following questions")
	channel := make(chan Attempt)
	timer := time.NewTimer(time.Duration(30) * time.Second)
outer:
	for index, quiz := range quizzes {
		go askQuestion(quiz, index, channel)
		select {
		case <-timer.C:
			attempts = append(attempts, Attempt{quiz: quiz, isTimedOut: true})
			break outer
		case attempt := <-channel:
			attempts = append(attempts, attempt)
		}
	}
	printScore(quizzes, attempts)
}
