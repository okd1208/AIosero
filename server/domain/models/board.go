package models

var ClientColor = 1
var ComputeColor = 2

type CellMatrix [8][8]int

type BoardStatus struct {
	Board                 CellMatrix
	ValidMovesForComputer [][]int
	ValidMovesForClient   [][]int
}

type NextMove struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
