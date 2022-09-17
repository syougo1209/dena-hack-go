package model

type EventGroup struct {
	EventID   EventID `json:"event_id" db:"event_id"`
	EventName string  `json:"event_name" db:"event_name"`
	Users     []*User `json:"users"`
}
