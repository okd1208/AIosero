package database

import (
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Game struct {
	ID            uint      `gorm:"primaryKey"`
	CreateAt      time.Time `gorm:"autoCreateTime"`
	isComputerWin bool
	BoardStates   []BoardState `gorm:"foreignKey:GameID"`
}

type BoardState struct {
	ID                   uint `gorm:"primaryKey"`
	GameID               uint
	Turn                 int
	Board                string
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
