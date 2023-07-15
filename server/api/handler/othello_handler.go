package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type OthelloRequest struct {
	UserColor string                      `json:"userColor"`
	Positions map[string]map[string][]int `json:"positions"`
}

type NextMove struct {
	row int
	col int
}

var clientColor = 1
var computeColor = 2

func evaluateMove(matrix [][]int, moveX int, moveY int, color int) int {
	// TODO: 評価関数を実装する
	return 0
}

func GetNextMovePosition(req OthelloRequest) (bestMove NextMove, err error) {
	matrix := positionsToMatrix(req.Positions, req.UserColor)

	// TODO:決定木に基づいてnextMoveを算出

	return bestMove, nil
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

func positionsToMatrix(positions map[string]map[string][]int, userColor string) [][]int {
	boardSize := 8
	matrix := make([][]int, boardSize)
	for i := range matrix {
		matrix[i] = make([]int, boardSize)
	}

	// クライアントを1、コンピュータ(自身)を2とする
	blackUser := computeColor
	whiteUser := computeColor

	if userColor == "black" {
		blackUser = clientColor
	} else if userColor == "white" {
		whiteUser = clientColor
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
