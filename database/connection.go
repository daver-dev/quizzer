package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	models "github.com/daver-dev/quizzer/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var questionsCollection *mongo.Collection
var questionsCollectionName = "questions"

var answersCollection *mongo.Collection
var answersCollectionName = "userAnswers"

func GetClient() *mongo.Client{
	uri := os.Getenv("DATABASE_URL")

	if client != nil {
	 return client
	}

	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx,options.Client().ApplyURI(uri))
	if err != nil {
	 log.Fatalln(err)
	}

	return client
}

// func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection{
// 	if questionsCollection != nil {
// 		return questionsCollection
// 	}
// 	questionsCollection := client.Database("quizzerDB").Collection(collectionName)
// 	return questionsCollection
// }

func GetCollection(client *mongo.Client, collectionName string, collection *mongo.Collection) *mongo.Collection{
	if collection != nil {
		return collection
	}
	newCollection := client.Database("quizzerDB").Collection(collectionName)
	return newCollection
}

func Disconnect(){
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	if client == nil{
	 return
	}
	err := client.Disconnect(ctx)
	if err != nil {
	 log.Fatalln(err)
	}
}

func Get_Questions()[]models.Question{
	client := GetClient()
	questionsCollection := GetCollection(client, questionsCollectionName, questionsCollection)
	ctx, cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()

	var questionsList []models.Question
	cursor, err := questionsCollection.Find(ctx,bson.D{})
	if err != nil {
	 log.Fatalln(err)
	 return nil
	}

	for cursor.Next(ctx){
	 var question models.Question
	 if err := cursor.Decode(&question); err != nil {
		log.Fatal(err)
	}
	 questionsList = append(questionsList, question)
	}
	
	return questionsList
}

func Search_Questions(searchString string)*models.Question{
	client := GetClient()
	questionsCollection := GetCollection(client, questionsCollectionName, questionsCollection)
	ctx, cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()

	var question *models.Question

	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: searchString}}}}
	err := questionsCollection.FindOne(ctx,filter).Decode(&question)
	if err != nil {
	return nil
	}
	return question
}

func Is_Answer_Correct(username string, questionId string, answerNumber int)bool {
	
	client := GetClient()
	questionsCollection := GetCollection(client, questionsCollectionName, questionsCollection)
	answersCollection := GetCollection(client, answersCollectionName, answersCollection)
	ctx, cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(questionId)
	filter := bson.D{{Key: "_id", Value: objID}}
	var question *models.Question

	questionsCollection.FindOne(ctx,filter).Decode(&question)
	// get the correct answer int
	fmt.Println(question.Correct_Answer)
	fmt.Printf("%#v", question)
	correctAnswerNumber := indexOfCorrectAnswer(question.Correct_Answer, question.Options) + 1
	if correctAnswerNumber == 0 {
		fmt.Println("Correct answer not found within answers")
	}
	fmt.Println(correctAnswerNumber)

	var answer = models.QuestionAnswer{
		Username: username,
		Question: question.Question,
		UserAnswer: question.Options[answerNumber - 1],
		WasCorrect: answerNumber == correctAnswerNumber}

	answersCollection.InsertOne(ctx, answer)

	return answer.WasCorrect
}

func indexOfCorrectAnswer(word string, data []string) (int) {
fmt.Println(word)
fmt.Println(data)

    for k, v := range data {
        if word == v {
            return k
        }
    }
    return -1
}
