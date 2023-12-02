package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/okd1208/OthelloLearning/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("not found")

type Game struct {
	ID            uint      `gorm:"primaryKey"`
	CreateAt      time.Time `gorm:"autoCreateTime"`
	IsComputerWin *bool     `gorm:"column:is_computer_win"`
	TotalTurn     int
	UserCount     int
	ComputerCount int
	BoardStates   []BoardState `gorm:"foreignKey:GameID"`
}

type BoardState struct {
	ID                   int `gorm:"primaryKey"`
	GameID               uint
	Turn                 int
	Board                string `gorm:"type:json"` // 二次元配列をJSON文字列として保存
	IsLastMoveByComputer bool
	LastMoveX            int
	LastMoveY            int
}

var DB *gorm.DB

func InitDB() error {
	var err error
	dsn := "host=postgres user=Othello password=password dbname=Othello port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// テーブルの自動マイグレーション
	err = DB.AutoMigrate(&Game{}, &BoardState{})
	if err != nil {
		return err
	}

	return nil
}

func CreateNewGame() (uint, error) {
	newGame := Game{
		IsComputerWin: nil,
	}

	if err := DB.Create(&newGame).Error; err != nil {
		return 0, err
	}
	fmt.Printf("(%%#v) %#v\n", &newGame)

	return newGame.ID, nil
}

func UpdateBoardState(gameID int, turn int, board models.CellMatrix, isLastMoveByComputer bool, lastMoveX, lastMoveY int) error {
	// boardをJSON文字列に変換する処理。encoding/jsonパッケージを使用。
	boardJSON, err := json.Marshal(board)
	if err != nil {
		return err
	}

	newBoardState := BoardState{
		GameID:               uint(gameID),
		Turn:                 turn,
		Board:                string(boardJSON),
		IsLastMoveByComputer: isLastMoveByComputer,
		LastMoveX:            lastMoveX,
		LastMoveY:            lastMoveY,
	}

	// データベースに保存
	if err := DB.Create(&newBoardState).Error; err != nil {
		return err
	}

	return nil
}

func GetBoardState(gameID int, turn int) (*BoardState, error) {
	var boardState BoardState
	err := DB.Where("game_id = ? AND turn = ?", uint(gameID), turn).First(&boardState).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// カスタムエラーを返す
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &boardState, nil
}

func UpdateGameResult(gameID, userCount, computerCount, turn int, isComputerWon bool) error {
	var game Game

	err := DB.First(&game, gameID).Error
	if err != nil {
		return err
	}

	game.IsComputerWin = &isComputerWon
	game.TotalTurn = turn
	game.UserCount = userCount
	game.ComputerCount = computerCount

	err = DB.Save(&game).Error
	if err != nil {
		return err
	}

	return nil
}
