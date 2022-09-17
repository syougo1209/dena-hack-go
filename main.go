package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/syougo1209/dena-hack-go/handler"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	mysqlDB, err := sql.Open("mysql", fmt.Sprintf("user:user@tcp(mysql:3306)/dena-hack?parseTime=true&loc=Local"))
	if err != nil {
		log.Printf("%v", err)
	}

	defer mysqlDB.Close()

	xdb := sqlx.NewDb(mysqlDB, "mysql")

	e := echo.New()
	uHandler := handler.UserHandler{Xdb: xdb}
	e.GET("/users/:id", uHandler.ServeHTTP)
	e.Logger.Fatal(e.Start(":8080"))
}

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
