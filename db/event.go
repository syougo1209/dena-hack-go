package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/dena-hack-go/model"
)

type EventRepository struct {
}

func (er *EventRepository) SelectByUserID(ctx context.Context, Db *sqlx.DB, id model.UserID) ([]*model.Event, error) {
	events := &[]*model.Event{}

	query := `SELECT e.id, e.name
	    FROM event as e
			join event_participation as ep on ep.event_id = e.id
			WHERE user_id = ?
	`

	if err := Db.SelectContext(ctx, events, query, id); err != nil {
		return nil, fmt.Errorf("GetContext user by id=%d: %w", id, err)
	}

	return *events, nil
}

func (er *EventRepository) GetByID(ctx context.Context, Db *sqlx.DB, id model.EventID) (*model.Event, error) {
	event := &model.Event{}

	query := `SELECT e.id, e.name
	    FROM event as e
			WHERE id = ?
	`

	if err := Db.GetContext(ctx, event, query, id); err != nil {
		return nil, fmt.Errorf("GetContext user by id=%d: %w", id, err)
	}

	return event, nil
}
