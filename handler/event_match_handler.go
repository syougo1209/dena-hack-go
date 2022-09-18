package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/dena-hack-go/db"
	"github.com/syougo1209/dena-hack-go/model"
)

type EventMatchHandler struct {
	Xdb *sqlx.DB
}
type EventMatchRequest struct {
	ID     int `param:"event_id"`
	UserID int `query:"user_id"`
}

func (uh *EventMatchHandler) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(EventMatchRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "params are inappropriate")
	}
	repository := db.UserRepository{}
	users, err := repository.SelectGroupUsersById(ctx, uh.Xdb, model.UserID(req.UserID))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	repository2 := db.EventRepository{}
	event, err := repository2.GetByID(ctx, uh.Xdb, model.EventID(req.ID))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	eg := model.EventGroup{
		EventID:   event.ID,
		EventName: event.Name,
		Users:     users,
	}

	return c.JSON(http.StatusOK, eg)
}
