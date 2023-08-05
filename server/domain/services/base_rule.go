package services

import (
	"github.com/okd1208/OthelloLearning/domain/models"
	"github.com/sirupsen/logrus"
)

func IsFinishGame(bs models.BoardStatus) bool {
	logger := logrus.New()
	if bs.CountEmptyCell == 0 {
		logger.Info("Game finished due to no empty cells left.")
		return true
	}
	if len(bs.ValidMovesForClient) == 0 && len(bs.ValidMovesForComputer) == 0 {
		logger.Info("Game finished as neither client nor computer has valid moves.")
		return true
	}
	return false
}
