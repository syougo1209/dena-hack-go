package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/dena-hack-go/model"
)

type UserRepository struct {
}

func (ur *UserRepository) FindById(ctx context.Context, Db *sqlx.DB, id model.UserID) (*model.User, error) {
	user := &model.User{}
	query := `SELECT *
	    FROM user
			WHERE id = ?
	`

	if err := Db.GetContext(ctx, user, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetContext user by id=%d: %w", id, err)
		}
		return nil, fmt.Errorf("GetContext user by id=%d: %w", id, err)
	}

	return user, nil
}
func (ur *UserRepository) SelectGroupUsersById(ctx context.Context, Db *sqlx.DB, id model.UserID, eid model.EventID) ([]*model.User, error) {
	dto := &DTO{}
	query := `SELECT egp.group_id, egp.user_id
	    FROM event_group_participation as egp
			JOIN event_group as eg on eg.id = egp.group_id
			WHERE egp.user_id = ?
			AND eg.event_id = ?
	`

	if err := Db.GetContext(ctx, dto, query, id, eid); err != nil {
		return nil, fmt.Errorf("GeContext user by id=%d: %w", id, err)
	}

	query2 := `SELECT u.id, u.name ,u.twitter_id
		FROM event_group_participation as egp
		JOIN user as u on u.id = egp.user_id
		WHERE egp.group_id = ?
	`

	users := &[]*model.User{}
	if err := Db.SelectContext(ctx, users, query2, dto.GroupID); err != nil {
		return nil, fmt.Errorf("GetContext user by id=%d: %w", id, err)
	}
	return *users, nil
}

type DTO struct {
	UserID  int `db:"user_id"`
	GroupID int `db:"group_id"`
}
