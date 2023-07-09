package main

import (
	"log"

	"github.com/okd1208/OthelloLearning/pkg/database"
	"github.com/okd1208/OthelloLearning/routes"

	// postgres ドライバ
	_ "github.com/lib/pq"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	routes.SetApi()
}
