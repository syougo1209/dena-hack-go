package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/dena-hack-go/db"
	"github.com/syougo1209/dena-hack-go/model"
)

type EventHandler struct {
	Xdb *sqlx.DB
}

type eventRequest struct {
	ID int `param:"event_id"`
}

func (eh *EventHandler) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(eventRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "params are inappropriate")
	}

	repository := &db.EventRepository{}
	event, err := repository.GetByID(ctx, eh.Xdb, model.EventID(req.ID))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, event)
}
