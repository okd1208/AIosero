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

type NextMove struct {
	Row int `json:"row"`
	Col int `json:"col"`
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

	validMoves := GetValidMoves(matrix)
	bestScore := -1
	for _, point := range validMoves {
		x := point[0]
		y := point[1]
		score := evaluateMove(matrix, x, y, computeColor)
		if score > bestScore {
			bestScore = score
			bestMove = NextMove{Row: x + 1, Col: y + 1} // フロントの仕様上一旦+1
		}
	}

	fmt.Println(bestMove)

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

func GetValidMoves(matrix [][]int) [][]int {
	validMoves := make([][]int, 0)
	directions := [8][2]int{
		{-1, -1}, // 左上
		{-1, 0},  // 上
		{-1, 1},  // 右上
		{0, -1},  // 左
		{0, 1},   // 右
		{1, -1},  // 左下
		{1, 0},   // 下
		{1, 1},   // 右下
	}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if matrix[x][y] != 0 {
				continue
			}
			for _, direction := range directions {
				dx, dy := direction[0], direction[1]
				if CheckMoveInDirection(matrix, x, y, dx, dy) {
					validMoves = append(validMoves, []int{x, y})
				}
			}
		}
	}
	return validMoves
}

// CheckMoveInDirectionは、与えられた方向(dx, dy)に対して石を置けるかどうかを探索する
// x, y は開始位置、dx, dyは探索方向です（-1, 0, 1のいずれか）
func CheckMoveInDirection(matrix [][]int, x, y, dx, dy int) bool {
	var isEnemyAdjacent bool // 敵の石が隣接しているかどうか
	x, y = x+dx, y+dy
	for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[x]) {
		if matrix[x][y] == 0 {
			return false
		}
		if matrix[x][y] == computeColor {
			return isEnemyAdjacent
		}
		if matrix[x][y] == clientColor && !isEnemyAdjacent {
			isEnemyAdjacent = true
		}
		x, y = x+dx, y+dy
	}
	return false
}
