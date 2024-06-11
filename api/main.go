package main

import (
	"log"

	"github.com/daver-dev/quizzer/database"
	"github.com/daver-dev/quizzer/routes"
	"github.com/daver-dev/quizzer/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var CommonHeaders = map[string]string{
	"Accept":                      "application/json",
	"Access-Control-Allow-Origin": "*",
	"X-Content-Type-Options":      "nosniff",
}

func main() {
	questions := utils.LoadQuestions("./assets/questions.json")
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Disconnect()
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		for key, value := range CommonHeaders {
			c.Header(key, value)
		}
		c.Next()
	})

	router.GET("/search", routes.SearchQuestions(questions))
	router.POST("/answer", routes.AnswerQuestion(questions))
	router.Run("localhost:8080")
}
