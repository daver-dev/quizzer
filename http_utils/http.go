package http_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/daver-dev/quizzer/models"
)

type AnswerResponse struct {
	Message string `json:"message"`
}

func HttpGetQuestion(searchString string) *models.Question {

	searchUrl := fmt.Sprintf("http://localhost:8080/search?searchString=%s", searchString)
	resp, err := http.Get(searchUrl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	// Unmarshal the JSON data into the struct
	var question models.Question
	err = json.Unmarshal(body, &question)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}
	return &question
}

func HttpPostAnswer(answerChoice string, questionId int) *string {
	questionIdString := fmt.Sprintf("%d", questionId)
	searchUrl := fmt.Sprintf("http://localhost:8080/answer?questionId=%s&answerNumber=%s", questionIdString, answerChoice)
	resp, err := http.Post(searchUrl, "application/json", nil)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	// Unmarshal the JSON data into the struct
	var result AnswerResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}
	resultString := fmt.Sprintf("%#v", result.Message)
	if resultString == `"Correct"` {
		fmt.Println("\nCorrect")
	} else {
		fmt.Println("\nIncorrect")
	}

	return &resultString
}
