package helpers

import "github.com/okd1208/OthelloLearning/domain/models"

func CountEmptyCell(board models.CellMatrix) int {
	count := 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				count++
			}
		}
	}
	return count
}

func GetReversedColor(color int) int {
	if color == 1 {
		return 2
	}
	return 1
}
