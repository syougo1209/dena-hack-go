package model

type EventID int64
type Event struct {
	ID   EventID `json:"id" db:"id"`
	Name string  `json:"name" db:"name"`
}
