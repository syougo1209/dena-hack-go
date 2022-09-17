package model

type UserID int64

type User struct {
	ID                 UserID   `json:"id" db:"id"`
	Name               string   `json:"name" db:"name"`
	Email              string   `json:"email" db:"email"`
	TwitterID          string   `json:"twitter_id" db:"twitter_id"`
	Password           string   `json:"password" db:"password"`
	ParticipatedEvents []*Event `json:"participated_events"`
}
