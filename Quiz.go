package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func askQuestion(id int, question string) {
	prefix := "Quiz question #" + fmt.Sprint(id)
	fmt.Print(prefix + ": " + question + "? ")
}

func readAndCheckAnswer(answer string) bool {
	var userAnswer string
	fmt.Scanf("%s", &userAnswer)
	return answer == userAnswer
}

func main() {
	quizFileFlag := flag.String("file", "problems.csv", "Fully qualified file path of the file containing the quizes")
	flag.Parse()
	file, err := os.Open(*quizFileFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var score, maxScore int
	for i, line := range lines {
		question, answer := line[0], line[1]
		askQuestion(i+1, question)
		if readAndCheckAnswer(answer) {
			score++
		}
		maxScore++
	}
	fmt.Println("Total user score: " + fmt.Sprint(score) + " out of " + fmt.Sprint(maxScore))
}
