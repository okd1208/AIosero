package services

import "github.com/okd1208/OthelloLearning/domain/models"

type Position struct {
	X int
	Y int
}

type PositionalWeight struct {
	Position Position
	Weight   int
}

var weights = []PositionalWeight{
	{Position: Position{0, 1}, Weight: -30},
	{Position: Position{1, 0}, Weight: -30},
	{Position: Position{1, 1}, Weight: -50},
	{Position: Position{0, 6}, Weight: -30},
	{Position: Position{1, 7}, Weight: -30},
	{Position: Position{1, 6}, Weight: -50},
	{Position: Position{7, 0}, Weight: -30},
	{Position: Position{7, 1}, Weight: -30},
	{Position: Position{6, 1}, Weight: -50},
	{Position: Position{7, 6}, Weight: -30},
	{Position: Position{6, 7}, Weight: -30},
	{Position: Position{6, 6}, Weight: -50},
}

func EvaluateMoveOption(board models.CellMatrix, x int, y int) int {
	nb, _ := GetNextBoard(board, x, y)
	clientValidMoves := GetValidMoves(nb, models.ClientColor)
	// 次相手の置くことのできるマスが少ないことを評価
	maxScore := 100
	leverage := 10
	score := maxScore - (len(clientValidMoves) * leverage)
	score = ApplyPositionalWeight(score, Position{x, y})
	return score
}

func ApplyPositionalWeight(score int, targetPos Position) int {
	for _, pw := range weights {
		if pw.Position == targetPos {
			return score + pw.Weight
		}
	}
	return score
}
