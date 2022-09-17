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
