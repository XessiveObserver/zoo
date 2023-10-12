package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// calls connection to the database method
	InitDB()

	// Calls endpoints method
	EndPoints()

	fmt.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
