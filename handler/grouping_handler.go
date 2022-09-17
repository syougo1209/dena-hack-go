package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/dena-hack-go/db"
	"github.com/syougo1209/dena-hack-go/model"
)

type GroupingHandler struct {
	Xdb *sqlx.DB
}

type groupingRequest struct {
	EventID int `param:"event_id"`
}

func (gh *GroupingHandler) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(groupingRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	repo := db.UserAnswerRepository{}
	choices, err := repo.MakeGroupingResponse(ctx, gh.Xdb, model.EventID(req.EventID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, choices)
}
