package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/syougo1209/dena-hack-go/db"
	"github.com/syougo1209/dena-hack-go/model"
	"github.com/syougo1209/dena-hack-go/service"
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

	// リクエスト用の構造体を用意する
	// [TODO]: グループ数を受け取れるようにする
	groupingRequest := model.GroupingRequest{
		GroupNum:     3,
		UsersChoices: choices,
	}

	// Pythonとのやりとりした結果を構造体で受け取る
	usersGroup, err := service.GroupingService(&groupingRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// 結果をデータベースに保存する

	return c.JSON(http.StatusOK, usersGroup)
}
