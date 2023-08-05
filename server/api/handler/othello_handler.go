package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/okd1208/OthelloLearning/domain/models"
	"github.com/okd1208/OthelloLearning/domain/services"
)

type OthelloRequest struct {
	UserColor string                      `json:"userColor"`
	Positions map[string]map[string][]int `json:"positions"`
}

type NextMove struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func HandleNextMove(c echo.Context) error {
	req := new(OthelloRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	matrix := positionsToMatrix(req.Positions, req.UserColor)

	nextMove, err := services.GetNextMovePosition(matrix)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nextMove)
}

func positionsToMatrix(positions map[string]map[string][]int, userColor string) models.CellMatrix {
	var matrix models.CellMatrix

	// クライアントを1、コンピュータ(自身)を2とする
	blackUser := models.ComputeColor
	whiteUser := models.ComputeColor

	if userColor == "black" {
		blackUser = models.ClientColor
	} else if userColor == "white" {
		whiteUser = models.ClientColor
	}

	for row, cols := range positions["black"] {
		rowInt, _ := strconv.Atoi(row)
		for _, col := range cols {
			matrix[rowInt-1][col-1] = blackUser
		}
	}
	for row, cols := range positions["white"] {
		rowInt, _ := strconv.Atoi(row)
		for _, col := range cols {
			matrix[rowInt-1][col-1] = whiteUser
		}
	}
	return matrix
}
