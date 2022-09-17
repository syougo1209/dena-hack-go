package model

type QuestionID int64
type Question struct {
	ID       QuestionID `json:"question_id" db:"id"`
	Question string     `json:"question" db:"content"`
	Answers  []Answer   `json:"answers"`
}

type Answer struct {
	QuestionID int    `db:"question_id"`
	Number     int    `json:"number" db:"number"`
	Content    string `json:"content" db:"content"`
}
