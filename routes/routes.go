package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daver-dev/quizzer/database"
	"github.com/daver-dev/quizzer/models"
	"github.com/daver-dev/quizzer/utils"
	"github.com/gin-gonic/gin"
)

// Route to search for questions with the query parameter searchString
func SearchQuestions(questions []models.Question) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchString := c.Query("searchString")
		question := utils.FindQuestion(searchString, questions)
		if question == nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No match for that question"})
			return
		}
		c.IndentedJSON(http.StatusOK, question)
	}
}

// Route to answer questions with the query parameters username, questionId, and answerNumber
func AnswerQuestion(questions []models.Question) gin.HandlerFunc {
	return func(c *gin.Context) {
		questionIdParam := c.Query("questionId")
		answerParam := c.Query("answerNumber")
		answerInt, err := strconv.Atoi(answerParam)
		if err != nil {
			fmt.Println("Can't convert answerNumber to an int.")
		}
		questionIdInt, err := strconv.Atoi(questionIdParam)
		if err != nil {
			fmt.Println("Can't convert questionId to an int.")
		}

		question := utils.GetQuestionById(questionIdInt, questions)
		isCorrect := utils.IsAnswerCorrect(*question, answerInt)

		var response string

		if isCorrect {
			response = "Correct"
		} else {
			response = "Incorrect"
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": response})

		var answer = models.Answer{Question: *question, UserAnswer: question.Options[answerInt-1]}
		database.SaveAnswer(answer)
	}
}
