package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// postgres ドライバ
	_ "github.com/lib/pq"
)

type combatExperience struct {
	ID        int
	IsUserWon string
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		dbTest()
		return c.String(http.StatusOK, "Hello, Echo World!!")
	})
	e.Logger.Fatal(e.Start(":8888"))
}

func dbTest() {
	var Db *sql.DB
	Db, err := sql.Open("postgres", "host=postgres user=Othello password=password dbname=Othello sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	sql := "SELECT id, Is_user_won FROM combat_experience WHERE id=$1;"

	pstatement, err := Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	queryID := 1
	var combatExperience combatExperience

	err = pstatement.QueryRow(queryID).Scan(&combatExperience.ID, &combatExperience.IsUserWon)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(combatExperience.ID, combatExperience.IsUserWon)
}
