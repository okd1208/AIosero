package main

import (
	"github.com/okd1208/OthelloLearning/routes"
	// postgres ドライバ
	_ "github.com/lib/pq"
)

func main() {
	routes.SetApi()
}
