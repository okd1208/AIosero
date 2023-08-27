package services

import (
	"errors"

	"github.com/okd1208/OthelloLearning/domain/helpers"
	"github.com/okd1208/OthelloLearning/domain/models"
)

func FullSearchStrategy(bs models.BoardStatus) ([]int, error) {
	// WARNING: CountEmptyCellが10以上は処理が一気に重くなる
	if bs.CountEmptyCell > 9 {
		return nil, errors.New("the number of empty cells is too large to process")
	}
	bestMove := []int{}
	bestScore := -100000

	for _, cell := range bs.ValidMovesForComputer {
		nb, err := GetNextBoard(bs.Board, cell[0], cell[1])
		if err != nil {
			return nil, err
		}

		newBS := models.BoardStatus{
			Board:                 nb,
			ValidMovesForComputer: GetValidMoves(nb, models.ComputeColor),
			ValidMovesForClient:   GetValidMoves(nb, models.ClientColor),
			CountEmptyCell:        helpers.CountEmptyCell(nb),
		}
		moveScore := minimax(newBS, 0, false)

		if moveScore > bestScore {
			bestScore = moveScore
			bestMove = []int{cell[0], cell[1]}
		}
	}

	return bestMove, nil
}

func minimax(bs models.BoardStatus, depth int, isComputerTurn bool) int {
	if IsFinishGame(bs) {
		// 一番下の階層でスコアを算出
		return calculateScore(bs)
	}

	if isComputerTurn {
		maxEval := -100000
		if IsSkipTurn(bs.ValidMovesForComputer) {
			eval := minimax(bs, depth, false)
			maxEval = max(maxEval, eval)
			return maxEval
		}
		for _, move := range bs.ValidMovesForComputer {
			nb, _ := GetNextBoard(bs.Board, move[0], move[1])
			newBS := models.BoardStatus{
				Board:                 nb,
				ValidMovesForComputer: GetValidMoves(nb, models.ComputeColor),
				ValidMovesForClient:   GetValidMoves(nb, models.ClientColor),
				CountEmptyCell:        helpers.CountEmptyCell(nb),
			}
			eval := minimax(newBS, depth+1, false)
			maxEval = max(maxEval, eval)
		}
		return maxEval
	} else {
		minEval := 100000
		if IsSkipTurn(bs.ValidMovesForClient) {
			eval := minimax(bs, depth, true)
			minEval = min(minEval, eval)
			return minEval
		}
		for _, move := range bs.ValidMovesForClient {
			nb, _ := GetNextBoard(bs.Board, move[0], move[1])
			newBS := models.BoardStatus{
				Board:                 nb,
				ValidMovesForComputer: GetValidMoves(nb, models.ComputeColor),
				ValidMovesForClient:   GetValidMoves(nb, models.ClientColor),
				CountEmptyCell:        helpers.CountEmptyCell(nb),
			}
			eval := minimax(newBS, depth+1, true)
			minEval = min(minEval, eval)
		}
		return minEval
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ゲームの状態を評価してスコアを計算する関数
func calculateScore(bs models.BoardStatus) int {
	score := 0
	for _, row := range bs.Board {
		for _, cell := range row {
			if cell == models.ComputeColor {
				score++
			} else if cell == models.ClientColor {
				score--
			}
		}
	}
	return score
}
