package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/dena-hack-go/model"
)

type GroupRepository struct {
}

func (ar *GroupRepository) StoreGroup(ctx context.Context, Db *sqlx.DB, id model.EventID, groups model.UsersGroup) error {
	userids := groups.Group

	query := `INSERT INTO event_group (event_id) VALUES (?);`
	query2 := `INSERT INTO event_group_participation (group_id, user_id) VALUES (?, ?);`

	for i, _ := range userids {
		r, err := Db.ExecContext(ctx, query, id)
		if err != nil {
			return fmt.Errorf("GetContext user by id=%d: %w", id, err)
		}
		group_id, err := r.LastInsertId()
		if err != nil {
			return fmt.Errorf("GetContext user by id=%d: %w", id, err)
		}

		for _, vj := range userids[i] {
			_, err := Db.ExecContext(ctx, query2, group_id, vj)
			if err != nil {
				return fmt.Errorf("GetContext user by id=%d: %w", id, err)
			}
		}
	}

	return nil
}
