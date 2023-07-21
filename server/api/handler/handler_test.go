package handler

import (
	"reflect"
	"testing"
)

func TestGetValidMoves(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "Initial state",
			matrix: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 2, 0, 0, 0},
				{0, 0, 0, 2, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: [][]int{{2, 3}, {3, 2}, {4, 5}, {5, 4}},
		},
		// More test cases...
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetValidMoves(tc.matrix)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("GetValidMoves() = %v, want %v", got, tc.expected)
			}
		})
	}
}
