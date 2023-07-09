package routes

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	api "github.com/okd1208/OthelloLearning/api/handler"
)

func SetApi() error {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/api/v1/othello/next-move", api.GetNextMovePosition)
	// e.GET("/api/v1/othello/next-move", func(c echo.Context) error {
	// 	userColor := c.QueryParam("userColor")
	// 	dbTest()
	// 	return c.String(http.StatusOK, "userColor = "+userColor)
	// })
	e.Logger.Fatal(e.Start(":8888"))
	return nil
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

type combatExperience struct {
	ID        int
	IsUserWon string
}
