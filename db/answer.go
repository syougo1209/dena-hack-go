package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/dena-hack-go/model"
)

type AnswerRepository struct {
}

func (ar *AnswerRepository) BulkInsertByUserID(ctx context.Context, Db *sqlx.DB, id model.UserID, answers []model.UserAnswer) error {

	query := `INSERT INTO user_answer (question_id, user_id, number) VALUES (?,?,?);`

	for _, v := range answers {
		if _, err := Db.ExecContext(ctx, query, v.QuestionID, id, v.Number); err != nil {
			return fmt.Errorf("GetContext user by id=%d: %w", id, err)
		}
	}

	return nil
}
