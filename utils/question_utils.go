package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	models "github.com/daver-dev/quizzer/models"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var answersCollection *mongo.Collection
var answersCollectionName = "userAnswers"

func GetClient() *mongo.Client {
	uri := os.Getenv("DATABASE_URL")

	if client != nil {
		return client
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func LoadQuestions(fileName string) []models.Question {
	var questions []models.Question
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	if file == nil {
		log.Fatalf("JSON file not found.")
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	err = json.Unmarshal(byteValue, &questions)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	// Assign each question an Id so the API can accept Ids for answer requests
	for i := range questions {
		questions[i].Id = i + 1
	}

	return questions
}

func FindQuestion(searchString string, questions []models.Question) *models.Question {
	for _, q := range questions {
		if containsIgnoreCase(q.Question, searchString) {
			return &q
		}
	}
	return nil
}

func containsIgnoreCase(stringToSearch, substr string) bool {
	lowerStringToSearch := strings.ToLower(stringToSearch)
	lowerSubstr := strings.ToLower(substr)
	return strings.Contains(lowerStringToSearch, lowerSubstr)
}

func GetQuestionById(id int, questions []models.Question) *models.Question {
	var foundQuestion models.Question
	for _, question := range questions {
		if question.Id == id {
			foundQuestion = question
		}
	}
	return &foundQuestion
}

func IsAnswerCorrect(question models.Question, answerNumber int) bool {
	// get the correct answer int
	correctAnswerNumber := indexOfCorrectAnswer(question.Correct_Answer, question.Options) + 1
	if correctAnswerNumber == 0 {
		fmt.Println("Correct answer not found within answers")
	}
	fmt.Println(correctAnswerNumber)

	return answerNumber == correctAnswerNumber
}

func indexOfCorrectAnswer(word string, data []string) int {
	fmt.Println(word)
	fmt.Println(data)

	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
