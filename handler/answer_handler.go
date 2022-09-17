package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/dena-hack-go/db"
	"github.com/syougo1209/dena-hack-go/model"
)

type AnswerHandler struct {
	Xdb *sqlx.DB
}

type answer struct {
	QuestionID int `json:"question_id"`
	Number     int `json:"choosed_number"`
}

type answerRequest struct {
	EventID int      `param:"event_id"`
	UserID  int      `json:"user_id"`
	Answers []answer `json:"answers"`
}

func (ah *AnswerHandler) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(answerRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	repository := &db.AnswerRepository{}
	answers := make([]model.UserAnswer, len(req.Answers))
	for i, v := range req.Answers {
		answers[i] = model.UserAnswer{
			QuestionID: v.QuestionID,
			Number:     v.Number,
		}
	}

	err := repository.BulkInsertByUserID(ctx, ah.Xdb, model.UserID(req.UserID), answers)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
