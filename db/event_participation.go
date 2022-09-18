package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/dena-hack-go/model"
)

type EventParticipationRepository struct {
}

func (er *EventParticipationRepository) Participate(ctx context.Context, Db *sqlx.DB, uid model.UserID, eid model.EventID) error {

	query := `INSERT INTO event_participation (user_id, event_id) VALUES (?,?);`
	if _, err := Db.ExecContext(ctx, query, uid, eid); err != nil {
		return fmt.Errorf("insert error %w", err)
	}
	return nil
}
