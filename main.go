package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/XessiveObserver/zoo/db"
	"github.com/XessiveObserver/zoo/routing"
)

func main() {
	// calls connection to the database method
	db.InitDB()

	// Calls endpoints method
	routing.EndPoints()

	fmt.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
