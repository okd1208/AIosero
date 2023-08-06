package services

import (
	"reflect"
	"testing"

	"github.com/okd1208/OthelloLearning/domain/helpers"
	"github.com/okd1208/OthelloLearning/domain/models"
)

func TestFullSearchStrategy(t *testing.T) {
	board := models.CellMatrix{
		{0, 2, 2, 2, 2, 2, 2, 0},
		{2, 1, 1, 1, 1, 1, 0, 1},
		{2, 1, 2, 1, 1, 1, 1, 1},
		{2, 1, 2, 1, 2, 2, 1, 1},
		{2, 1, 0, 2, 1, 2, 1, 1},
		{2, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 0, 1, 1},
		{0, 2, 2, 0, 2, 2, 2, 2},
	}

	bs := models.BoardStatus{
		Board:                 board,
		ValidMovesForComputer: GetValidMoves(board, models.ComputeColor),
		ValidMovesForClient:   GetValidMoves(board, models.ClientColor),
		CountEmptyCell:        helpers.CountEmptyCell(board), // 残りの空きマスの数
	}

	bestMove, err := FullSearchStrategy(bs)
	if err != nil {
		t.Fatalf("Failed to compute best move: %v", err)
	}

	expectedBestMove := []int{0, 0}

	if !reflect.DeepEqual(expectedBestMove, bestMove) {
		t.Errorf("Expected best move to be %v but got %v", expectedBestMove, bestMove)
	}
}
