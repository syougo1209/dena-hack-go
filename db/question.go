package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/dena-hack-go/model"
)

type QuestionRepository struct {
}

func (ur *QuestionRepository) SelectAll(ctx context.Context, Db *sqlx.DB) ([]*model.Question, error) {
	questions := &[]*model.Question{}
	query := `SELECT *
	    FROM question
	`

	if err := Db.SelectContext(ctx, questions, query); err != nil {
		return nil, fmt.Errorf("GetContext user: %w", err)
	}
	ids := make([]model.QuestionID, len(*questions))
	for i, v := range *questions {
		ids[i] = v.ID
	}

	answers := &[]*model.Answer{}
	query2 := `SELECT *
	    FROM answer
			Where question_id in (?)
			Order by question_id
	`
	sql, params, err := sqlx.In(query2, ids)
	if err != nil {
		return nil, fmt.Errorf("GetContext user: %w", err)
	}
	query3 := Db.Rebind(sql)

	if err := Db.Select(answers, query3, params...); err != nil {
		log.Fatal(err)
	}
	for _, av := range *answers {
		for _, qv := range *questions {
			if int(qv.ID) == av.QuestionID {
				qv.Answers = append(qv.Answers, *av)
			}
		}
	}

	return *questions, nil
}
