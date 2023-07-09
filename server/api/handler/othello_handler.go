package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type OthelloRequest struct {
	UserColor string                      `json:"userColor"`
	Positions map[string]map[string][]int `json:"positions"`
}

func GetNextMovePosition(req OthelloRequest) (nextMove OthelloRequest, err error) {
	matrix := positionsToMatrix(req.Positions)
	fmt.Println(matrix)

	// TODO:深層学習に基づいてnextMoveを算出
	nextMove = req

	return nextMove, nil
}

func HandleNextMove(c echo.Context) error {
	req := new(OthelloRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	nextMove, err := GetNextMovePosition(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nextMove)
}

func positionsToMatrix(positions map[string]map[string][]int) [][]int {
	boardSize := 8
	matrix := make([][]int, boardSize)
	for i := range matrix {
		matrix[i] = make([]int, boardSize)
	}

	// 黒の駒を2、白の駒を1
	for row, cols := range positions["black"] {
		rowInt, _ := strconv.Atoi(row)
		for _, col := range cols {
			matrix[rowInt-1][col-1] = 2
		}
	}
	for row, cols := range positions["white"] {
		rowInt, _ := strconv.Atoi(row)
		for _, col := range cols {
			matrix[rowInt-1][col-1] = 1
		}
	}
	return matrix
}
