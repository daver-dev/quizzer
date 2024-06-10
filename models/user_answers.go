package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

//better models, but not enough time to implement

// type QuestionAnswer struct {
// 	Question string `json:"question"`
// 	UserAnswer string `json:"user_answer"`
// 	WasCorrect bool `json:"was_correct"`
// }

// type UserAnswers struct {
// 	Id    primitive.ObjectID   `json:"id"`
// 	Username  string      `json:"username"`
// 	Answers []QuestionAnswer `json:"answers"`
// }

type QuestionAnswer struct {
	Username string `json:"username"`
	Question string `json:"question"`
	UserAnswer string `json:"user_answer"`
	WasCorrect bool `json:"was_correct"`
}