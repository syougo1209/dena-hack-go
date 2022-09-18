package handler

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type PrepareDataHandler struct {
	Xdb *sqlx.DB
}

func (ph *PrepareDataHandler) ServeHTTP(c echo.Context) error {
	ctx := c.Request().Context()
	query := `INSERT INTO user (name, twitter_id, email, password) VALUES (?,?,?, ?);`

	var uids []int64
	for i := 0; i < 5; i++ {
		r, err := ph.Xdb.ExecContext(ctx, query, fmt.Sprintf("name%d", i), "twitter_id", "email", "password")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		id, _ := r.LastInsertId()
		uids = append(uids, id)
	}

	query2 := `INSERT INTO event (name) VALUES (?);`
	r, _ := ph.Xdb.ExecContext(ctx, query2, "就活イベント")
	eventID, _ := r.LastInsertId()
	query3 := `INSERT INTO event_participation (user_id, event_id) VALUES (?,?);`
	for _, v := range uids {
		_, err := ph.Xdb.ExecContext(ctx, query3, v, eventID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	//question1
	query4 := `INSERT INTO question (content) VALUES (?);`
	q, err := ph.Xdb.ExecContext(ctx, query4, "好きな料理は何ですか？")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id, _ := q.LastInsertId()
	query5 := `INSERT INTO answer (question_id, content, number) VALUES (?,?,?);`
	r, err = ph.Xdb.ExecContext(ctx, query5, id, "中華料理", 1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	r, err = ph.Xdb.ExecContext(ctx, query5, id, "フランス料理", 2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	r, err = ph.Xdb.ExecContext(ctx, query5, id, "日本料理", 3)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//question2
	q, err = ph.Xdb.ExecContext(ctx, query4, "好きなフルーツは何ですか？")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id, _ = q.LastInsertId()
	r, err = ph.Xdb.ExecContext(ctx, query5, id, "ぶどう", 1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	r, err = ph.Xdb.ExecContext(ctx, query5, id, "りんご", 2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	r, err = ph.Xdb.ExecContext(ctx, query5, id, "なし", 3)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//question3
	q, err = ph.Xdb.ExecContext(ctx, query4, "ゲームは好きですか？")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id, _ = q.LastInsertId()
	r, err = ph.Xdb.ExecContext(ctx, query5, id, "はい", 1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	r, err = ph.Xdb.ExecContext(ctx, query5, id, "いいえ", 2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
