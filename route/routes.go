package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func EndPoints() {
	router := httprouter.New()
	router.GET("/animals", GetAnimals)
	router.GET("/animals/:id", GetAnimal)
	router.POST("/animals", CreateAnimal)
	router.PUT("/animals/:id", UpdateAnimal)
	router.DELETE("/animals/:id", DeleteAnimal)

	fmt.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
