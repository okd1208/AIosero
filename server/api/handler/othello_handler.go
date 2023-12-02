package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/okd1208/OthelloLearning/domain/helpers"
	"github.com/okd1208/OthelloLearning/domain/models"
	"github.com/okd1208/OthelloLearning/domain/services"
	"github.com/okd1208/OthelloLearning/internal/database"
)

type OthelloRequest struct {
	UserColor string                      `json:"userColor"`
	Positions map[string]map[string][]int `json:"positions"`
	GameID    int                         `json:"gameId"`
	Turn      int                         `json:"turn"`
	LastPut   map[string]int              `json:"lastPut"`
}

type NextMove struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func HandleNextMove(c echo.Context) error {
	req := new(OthelloRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	matrix := positionsToMatrix(req.Positions, req.UserColor)

	lastPut := map[string]int{
		"row": req.LastPut["row"] - 1,
		"col": req.LastPut["col"] - 1,
	}

	database.UpdateBoardState(req.GameID, req.Turn, matrix, true, lastPut["row"], lastPut["col"])
	nextMove, err := services.GetNextMovePosition(matrix)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	newMatrix, err := services.GetNextBoard(matrix, nextMove.Row, nextMove.Col)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	database.UpdateBoardState(req.GameID, req.Turn+1, newMatrix, false, nextMove.Row, nextMove.Col)
	nextMove.Row = nextMove.Row + 1 // フロントの仕様上一旦+1
	nextMove.Col = nextMove.Col + 1
	return c.JSON(http.StatusOK, nextMove)
}

func CollectFinishGameData(c echo.Context) error {
	req := new(OthelloRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	matrix := positionsToMatrix(req.Positions, req.UserColor)

	// 最新のBoardStateが登録されているかを確認し、されていなければ登録
	_, err := database.GetBoardState(req.GameID, req.Turn)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			lastPut := map[string]int{
				"row": req.LastPut["row"] - 1,
				"col": req.LastPut["col"] - 1,
			}
			database.UpdateBoardState(req.GameID, req.Turn, matrix, true, lastPut["row"], lastPut["col"])
		} else {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}
	userCount := helpers.CountClientCell(matrix)
	computerCount := helpers.CountComputerCell(matrix)
	isComputerWon := computerCount > userCount

	err = database.UpdateGameResult(req.GameID, userCount, computerCount, req.Turn, isComputerWon)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func positionsToMatrix(positions map[string]map[string][]int, userColor string) models.CellMatrix {
	var matrix models.CellMatrix

	// クライアントを1、コンピュータ(自身)を2とする
	blackUser := models.ComputeColor
	whiteUser := models.ComputeColor

	if userColor == "black" {
		blackUser = models.ClientColor
	} else if userColor == "white" {
		whiteUser = models.ClientColor
	}

	for row, cols := range positions["black"] {
		rowInt, _ := strconv.Atoi(row)
		for _, col := range cols {
			matrix[rowInt-1][col-1] = blackUser
		}
	}
	for row, cols := range positions["white"] {
		rowInt, _ := strconv.Atoi(row)
		for _, col := range cols {
			matrix[rowInt-1][col-1] = whiteUser
		}
	}
	return matrix
}
