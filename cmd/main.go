package main

import (
	"github.com/Food-to-Share/server/database"
	"github.com/Food-to-Share/server/server"
	_ "github.com/lib/pq"
)

func main() {
	database.StartDB()

	// server := server.NewServer()

	server.Run()
}
