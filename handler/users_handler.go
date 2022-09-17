package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/dena-hack-go/db"
	"github.com/syougo1209/dena-hack-go/model"
)

type UserHandler struct {
	Xdb *sqlx.DB
}
type userRequest struct {
	ID int `param:"id"`
}

func (uh *UserHandler) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(userRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "params are inappropriate")
	}
	repository := &db.UserRepository{}
	user, err := repository.FindById(ctx, uh.Xdb, model.UserID(req.ID))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	repository2 := &db.EventRepository{}
	events, err := repository2.SelectByUserID(ctx, uh.Xdb, model.UserID(req.ID))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	user.ParticipatedEvents = events
	return c.JSON(http.StatusOK, user)
}
