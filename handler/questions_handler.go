package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/dena-hack-go/db"
)

type QuestionsHandler struct {
	Xdb *sqlx.DB
}
type questionsRequest struct {
	EventID int `param:"event_id"`
}

func (qh *QuestionsHandler) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(questionsRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "params are inappropriate")
	}
	repository := &db.QuestionRepository{}
	questions, err := repository.SelectAll(ctx, qh.Xdb)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, questions)
}
