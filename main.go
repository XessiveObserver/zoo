package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/XessiveObserver/zoo/db"
)

func main() {
	// calls connection to the database method
	db.InitDB()

	fmt.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
