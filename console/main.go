package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/daver-dev/quizzer/http_utils"
	"github.com/daver-dev/quizzer/utils"
)

func main() {
	utils.LoadQuestions("./assets/questions.json")

	
	var questionToAsk string
	var userAnswer string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=============Quizzer=============")
	for {
		fmt.Print("\nSearch for a question topic: ")
		var searchTopic string
		searchTopic, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		
		searchTopic = strings.TrimSpace(searchTopic)

		if searchTopic == "" {
			continue
		}

		retreivedQuestion := http_utils.HttpGetQuestion(searchTopic)

		if retreivedQuestion.Id != 0 {
			questionToAsk = fmt.Sprintf("\nQuestion: %s\n\nOptions:\n\n1. %s\n\n2. %s\n\n3. %s\n\n4. %s\n", retreivedQuestion.Question, retreivedQuestion.Options[0], retreivedQuestion.Options[1], retreivedQuestion.Options[2], retreivedQuestion.Options[3])
			fmt.Println(questionToAsk)
			fmt.Println("To answer, enter 1 through 4")
			fmt.Println()
			fmt.Scan(&userAnswer)
			http_utils.HttpPostAnswer(userAnswer, retreivedQuestion.Id)
		} else {
			fmt.Println("\nThere are no questions including that topic.")
		}
	}
}
