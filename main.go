package main

import (
	"log"

	handlers "github.com/avalith-net/skill-factory-go-carrerpath/controller"
	db "github.com/avalith-net/skill-factory-go-carrerpath/database"
)

func main() {

	if db.ConnectionCheck() == 0 {
		log.Fatal("No connection to the database")
		return
	}
	handlers.LaunchingServer()
}
