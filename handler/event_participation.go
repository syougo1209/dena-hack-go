package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/dena-hack-go/db"
	"github.com/syougo1209/dena-hack-go/model"
)

type EventParticipationHandler struct {
	Xdb *sqlx.DB
}

type participationReq struct {
	ID     int `param:"event_id"`
	UserID int `json:"user_id"`
}

func (eph *EventParticipationHandler) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(participationReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	repository := &db.EventParticipationRepository{}
	err := repository.Participate(ctx, eph.Xdb, model.UserID(req.UserID), model.EventID(req.ID))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
