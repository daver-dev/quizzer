package models

type Question struct {
	Id             int      `json:"id"`
	Question       string   `json:"question"`
	Options        []string `json:"options"`
	Correct_Answer string   `json:"correct_answer"`
	Distractors    []string `json:"distractors"`
}

type Answer struct {
	Question   Question `json:"question"`
	UserAnswer string   `json:"user_answer"`
}
