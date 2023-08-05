package utils // utilsパッケージとします

import (
	"fmt"
)

// PrintBoard is a debug utility function to print the board in a readable format
func PrintBoard(board [8][8]int) {
	log := "--------\n"
	for _, row := range board {
		for _, cell := range row {
			log += fmt.Sprintf("%d ", cell)
		}
		log += "\n"
	}
	fmt.Println(log)
}

func colorRed(text string) string {
	return "\033[31m" + text + "\033[0m"
}

func CompareAndPrintBoards(board1, board2 [8][8]int) {
	log := "--------\n"
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board1[i][j] != board2[i][j] {
				log += fmt.Sprintf("%s ", colorRed(fmt.Sprint(board2[i][j])))
			} else {
				log += fmt.Sprintf("%d ", board2[i][j])
			}
		}
		log += "\n"
	}
	fmt.Println(log)
}
