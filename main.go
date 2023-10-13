package main

import (
	"github.com/XessiveObserver/zoo/db"
	"github.com/XessiveObserver/zoo/route"
)

func main() {
	// calls connection to the database method
	db.InitDB()

	// Calls endpoints method
	route.EndPoints()
}
