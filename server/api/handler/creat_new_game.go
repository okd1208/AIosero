package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/okd1208/OthelloLearning/internal/database"
)

func CreateNewGame(c echo.Context) error {
	gameId, err := database.CreateNewGame()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, gameId)
}
