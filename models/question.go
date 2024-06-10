package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Id    primitive.ObjectID	`bson:"_id,omitempty" json:"id,omitempty"`
	Question  string      `json:"question"`
	Options  []string      `json:"options"`
	Correct_Answer string `json:"correct_answer"`
	Distractors []string `json:"distractors"`
   }