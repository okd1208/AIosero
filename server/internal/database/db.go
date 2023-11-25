package database

import (
	"encoding/json"
	"time"

	_ "github.com/lib/pq"
	"github.com/okd1208/OthelloLearning/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Game struct {
	ID            uint         `gorm:"primaryKey"`
	CreateAt      time.Time    `gorm:"autoCreateTime"`         // CreateAt を CreatedAt に変更
	IsComputerWin *bool        `gorm:"column:is_computer_win"` // 最初の文字を大文字に変更
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
		CreateAt:      time.Now(),
	}

	if err := DB.Create(&newGame).Error; err != nil {
		return 0, err
	}

	return newGame.ID, nil
}

func UpdateBoardState(gameID int, turn int, board models.CellMatrix, isLastMoveByComputer bool, lastMoveX, lastMoveY int) error {
	// boardをJSON文字列に変換する処理。encoding/jsonパッケージを使用。
	boardJSON, err := json.Marshal(board)
	if err != nil {
		return err
	}

	// 新しいBoardStateを作成
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
