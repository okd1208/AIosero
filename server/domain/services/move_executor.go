package services

import (
	"fmt"

	"github.com/okd1208/OthelloLearning/domain/helpers"
	"github.com/okd1208/OthelloLearning/domain/models"

	"github.com/sirupsen/logrus"
)

var directions = [8][2]int{
	{-1, -1}, // 左上
	{-1, 0},  // 上
	{-1, 1},  // 右上
	{0, -1},  // 左
	{0, 1},   // 右
	{1, -1},  // 左下
	{1, 0},   // 下
	{1, 1},   // 右下
}

func GetNextMovePosition(board models.CellMatrix) (bestMove models.NextMove, err error) {
	logger := logrus.New()
	bs := models.BoardStatus{
		Board:                 board,
		ValidMovesForComputer: GetValidMoves(board, models.ComputeColor),
		ValidMovesForClient:   GetValidMoves(board, models.ClientColor),
		CountEmptyCell:        helpers.CountEmptyCell(board),
	}
	if IsSkipTurn(bs.ValidMovesForComputer) {
		logger.Info("computer's turn is skip")
		return models.NextMove{Row: -1, Col: -1}, nil
	}

	if bs.CountEmptyCell < 9 {
		bestMove, err := FullSearchStrategy(bs)
		if err != nil {
			return models.NextMove{Row: -1, Col: -1}, err
		}
		return models.NextMove{Row: bestMove[0] + 1, Col: bestMove[1] + 1}, nil
	}
	bestScore := -1000
	for _, point := range bs.ValidMovesForComputer {
		x := point[0]
		y := point[1]
		score := EvaluateMoveOption(board, x, y)
		if score > bestScore {
			bestScore = score
			bestMove = models.NextMove{Row: x + 1, Col: y + 1} // フロントの仕様上一旦+1
		}
	}

	fmt.Println(bestMove)

	return bestMove, nil
}

func GetValidMoves(board models.CellMatrix, myColor int) [][]int {
	validMoves := make([][]int, 0)

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] != 0 {
				continue
			}
			for _, direction := range directions {
				dx, dy := direction[0], direction[1]
				if CheckMoveInDirection(board, x, y, dx, dy, myColor) {
					validMoves = append(validMoves, []int{x, y})
				}
			}
		}
	}
	return validMoves
}

// CheckMoveInDirectionは、与えられた方向(dx, dy)に対して石を置けるかどうかを探索する
// x, y は開始位置、dx, dyは探索方向です（-1, 0, 1のいずれか）
func CheckMoveInDirection(board models.CellMatrix, x, y, dx, dy int, myColor int) bool {
	var isEnemyAdjacent bool // 敵の石が隣接しているかどうか
	enemyColor := helpers.GetReversedColor(myColor)
	x, y = x+dx, y+dy
	for x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) {
		if board[x][y] == 0 {
			return false
		}
		if board[x][y] == myColor {
			return isEnemyAdjacent
		}
		if board[x][y] == enemyColor && !isEnemyAdjacent {
			isEnemyAdjacent = true
		}
		x, y = x+dx, y+dy
	}
	return false
}

func GetNextBoard(board models.CellMatrix, x int, y int) (reversedBoard models.CellMatrix, err error) {
	for _, direction := range directions {
		dx, dy := direction[0], direction[1]
		if CheckMoveInDirection(board, x, y, dx, dy, models.ComputeColor) {
			reversedBoard, err = reverseCell(board, x, y, dx, dy)
			if err != nil {
				return reversedBoard, err
			}
		}
	}
	return reversedBoard, err
}

func reverseCell(board models.CellMatrix, x, y, dx, dy int) (models.CellMatrix, error) {
	board[x][y] = models.ComputeColor
	x, y = x+dx, y+dy
	for x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) {
		if board[x][y] == models.ComputeColor {
			return board, nil
		}
		if board[x][y] == models.ClientColor {
			board[x][y] = models.ComputeColor
		}
		x, y = x+dx, y+dy
	}
	return board, fmt.Errorf("can't put a cell at x: %d, y: %d", x, y)
}

func evaluateMove(matrix models.CellMatrix, moveX int, moveY int) int {
	// TODO: 評価関数を実装する
	return 0
}
