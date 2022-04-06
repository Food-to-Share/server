package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Food-to-Share/server/servers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "6910"
	dbname   = "foodToShare"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	servers.Start()

	fmt.Println("Server Started!")
}

