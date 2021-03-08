package main

import (
	"log"

	"github.com/avalith-net/skill-factory-go-carrerpath/controller"
	"github.com/avalith-net/skill-factory-go-carrerpath/database"
)

func main() {

	if database.ConnectionCheck() == 0 {
		log.Fatal("No connection to the database")
		return
	}
	controller.LaunchingServer()
}
