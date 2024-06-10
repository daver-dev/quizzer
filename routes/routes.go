package routes

import (
	"fmt"
	"net/http"
	"strconv"

	database "github.com/daver-dev/quizzer/database"
	"github.com/gin-gonic/gin"
)

func ListQuestions(c *gin.Context) {
	questionsList := database.Get_Questions()
	if questionsList == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"There are no questions"})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "Successfully retrieved questions", "data":questionsList})
}

func SearchQuestions(c *gin.Context) {
	searchString := c.Query("searchString")
	question := database.Search_Questions(searchString)
	if question == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"No match for that question"})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "Successfully retrieved questions", "data":question})
}

func AnswerQuestion(c *gin.Context) {
	username := c.Query("username")
	questionId := c.Query("questionId")
	answerNumber := c.Query("answerNumber")


	answerInt, err := strconv.Atoi(answerNumber);
	if err != nil {
		fmt.Println("Can't convert this to an int!")
	}
	
	isCorrect := database.Is_Answer_Correct(username, questionId, answerInt)
	var response string

	if isCorrect {
		response = "Correct"
	} else {
		response = "Incorrect"
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": response})
}