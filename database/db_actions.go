package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/daver-dev/quizzer/models"
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

func GetCollection(client *mongo.Client, collectionName string, collection *mongo.Collection) *mongo.Collection {
	if collection != nil {
		return collection
	}
	newCollection := client.Database("quizzerDB").Collection(collectionName)
	return newCollection
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if client == nil {
		return
	}
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func SaveAnswer(answer models.Answer) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := GetClient()
	answersCollection := GetCollection(client, answersCollectionName, answersCollection)
	answersCollection.InsertOne(ctx, answer)
}
