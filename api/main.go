package main

import (
	"fmt"
	"log"

	database "github.com/daver-dev/quizzer/database"
	routes "github.com/daver-dev/quizzer/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() { 
	err := godotenv.Load()
  if err != nil {
   log.Fatalln(err)
  }
	defer database.Disconnect()
	router := gin.Default()
	router.GET("/questions", routes.ListQuestions)
	router.GET("/search", routes.SearchQuestions)
	router.POST("/answer", routes.AnswerQuestion)
	router.Run("localhost:8080")
	fmt.Println("yo")
}

// 6665ba7b0591bb5f9ece7629