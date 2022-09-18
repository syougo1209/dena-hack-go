package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/dena-hack-go/model"
)

type UserAnswerRepository struct {
}

func (uar *UserAnswerRepository) MakeGroupingResponse(ctx context.Context, Db *sqlx.DB, id model.EventID) ([]model.UsersChoice, error) {
	query := `SELECT ep.user_id
		FROM event_participation as ep
		WHERE ep.event_id = ?
	`
	users := &[]dto{}
	if err := Db.SelectContext(ctx, users, query, id); err != nil {
		return nil, fmt.Errorf("GetContext user by id=%d: %w", id, err)
	}

	req := make([]model.UsersChoice, len(*users))

	query2 := `SELECT *
		FROM user_answer as ea
		WHERE ea.user_id = ?
	`

	for i, user := range *users {
		req[i].UserID = int(user.ID)
		qanswers := &[]uadto{}
		if err := Db.SelectContext(ctx, qanswers, query2, user.ID); err != nil {
			return nil, fmt.Errorf("GetContext user by id=%d: %w", id, err)
		}
		choices := make([]int, len(*qanswers))
		for i, qa := range *qanswers {
			choices[i] = qa.Number
		}
		req[i].Choices = choices
	}
	return req, nil
}

type dto struct {
	ID int `db:"user_id"`
}
type uadto struct {
	QuestionID int `db:"question_id"`
	UserID     int `db:"user_id"`
	Number     int `db:"number"`
}
