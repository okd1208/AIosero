package game

import (
	"fmt"

	"github.com/okd1208/OthelloLearning/domain/models"
)

func GetNextMovePosition(board models.Board) (bestMove models.NextMove, err error) {
	validMoves := GetValidMoves(board.Cells)
	bestScore := -1
	for _, point := range validMoves {
		x := point[0]
		y := point[1]
		score := evaluateMove(board.Cells, x, y)
		if score > bestScore {
			bestScore = score
			bestMove = models.NextMove{Row: x + 1, Col: y + 1} // フロントの仕様上一旦+1
		}
	}

	fmt.Println(bestMove)

	return bestMove, nil
}

func GetValidMoves(board models.CellMatrix) [][]int {
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

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] != 0 {
				continue
			}
			for _, direction := range directions {
				dx, dy := direction[0], direction[1]
				if CheckMoveInDirection(board, x, y, dx, dy) {
					validMoves = append(validMoves, []int{x, y})
				}
			}
		}
	}
	return validMoves
}

// CheckMoveInDirectionは、与えられた方向(dx, dy)に対して石を置けるかどうかを探索する
// x, y は開始位置、dx, dyは探索方向です（-1, 0, 1のいずれか）
func CheckMoveInDirection(board models.CellMatrix, x, y, dx, dy int) bool {
	var isEnemyAdjacent bool // 敵の石が隣接しているかどうか
	x, y = x+dx, y+dy
	for x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) {
		if board[x][y] == 0 {
			return false
		}
		if board[x][y] == models.ComputeColor {
			return isEnemyAdjacent
		}
		if board[x][y] == models.ClientColor && !isEnemyAdjacent {
			isEnemyAdjacent = true
		}
		x, y = x+dx, y+dy
	}
	return false
}

func evaluateMove(matrix models.CellMatrix, moveX int, moveY int) int {
	// TODO: 評価関数を実装する
	return 0
}
