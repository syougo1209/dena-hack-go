package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

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
	e.GET("/", func(c echo.Context) error {
		ctx := c.Request().Context()
		query := `INSERT user (name, email, password) VALUES (?,?,?)`
		user := &User{Name: "name", Email: "email", Password: "password"}
		xdb.ExecContext(ctx, query, user.Name, user.Email, user.Password)
		return c.String(http.StatusOK, "Hello, Worlds!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
