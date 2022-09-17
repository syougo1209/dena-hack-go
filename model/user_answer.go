package model

type UserAnswer struct {
	QuestionID int `db:"question_id"`
	Number     int `json:"number" db:"number"`
}

type UserAnswers []UserAnswer
