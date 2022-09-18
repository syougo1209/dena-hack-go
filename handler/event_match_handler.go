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

	userreses := make([]userres, len(users))

	for i, v := range users {
		userreses[i] = userres{
			ID:        int(v.ID),
			Name:      v.Name,
			TwitterID: v.TwitterID,
		}
	}

	eg := response{
		EventID:   int(event.ID),
		EventName: event.Name,
		Users:     userreses,
	}

	return c.JSON(http.StatusOK, eg)
}

type response struct {
	EventID   int       `json:"event_id"`
	EventName string    `json:"event_name"`
	Users     []userres `json:"users"`
}

type userres struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TwitterID string `json:"twitter_id"`
}
