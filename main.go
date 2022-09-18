package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/syougo1209/dena-hack-go/config"
	"github.com/syougo1209/dena-hack-go/handler"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Printf("failed to new config: %v", err)
	}

	mysqlDB, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true&loc=Local",
		cfg.MySQLUser, cfg.MYSQLPassword,
		cfg.MYSQLAddr, cfg.MYSQLDbName,
	))

	if err != nil {
		log.Printf("%v", err)
	}

	defer mysqlDB.Close()

	xdb := sqlx.NewDb(mysqlDB, "mysql")

	log.Printf("hello!!world!!")
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	uHandler := handler.UserHandler{Xdb: xdb}
	e.GET("/users/:id", uHandler.ServeHTTP)

	emHandler := handler.EventMatchHandler{Xdb: xdb}
	e.GET("/events/:event_id/matching", emHandler.ServeHTTP)

	qHandler := handler.QuestionsHandler{Xdb: xdb}
	e.GET("/events/:event_id/questions", qHandler.ServeHTTP)

	aHandler := handler.AnswerHandler{Xdb: xdb}
	e.POST("/events/:event_id/questions", aHandler.ServeHTTP)

	gHandler := handler.GroupingHandler{Xdb: xdb}
	e.GET("/events/:event_id/grouping", gHandler.ServeHTTP)

	pHandler := handler.PrepareDataHandler{Xdb: xdb}
	e.POST("/prepare_data", pHandler.ServeHTTP)
	e.Logger.Fatal(e.Start(":80"))
}
