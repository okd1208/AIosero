package helpers

import (
	"testing"

	"github.com/okd1208/OthelloLearning/domain/models"
)

func TestCountEmptyCell(t *testing.T) {
	board := models.CellMatrix{
		{0, 0, 1, 2, 0, 0, 0, 0},
		{0, 1, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	count := CountEmptyCell(board)

	if count != 60 {
		t.Errorf("Count was incorrect, got: %d, want: %d.", count, 60)
	}
}

func TestGetReversedColor(t *testing.T) {
	color := 1
	expectedColor := 2

	resultColor := GetReversedColor(color)

	if resultColor != expectedColor {
		t.Errorf("Reversed color was incorrect, got: %d, want: %d.", resultColor, expectedColor)
	}

	color = 2
	expectedColor = 1

	resultColor = GetReversedColor(color)

	if resultColor != expectedColor {
		t.Errorf("Reversed color was incorrect, got: %d, want: %d.", resultColor, expectedColor)
	}
}
