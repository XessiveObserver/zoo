package routing

import (
	"github.com/julienschmidt/httprouter"
)

func EndPoints() {
	router := httprouter.New()
	router.GET("/animals", GetAnimals)
	router.GET("/animals:id", GetAnimal)
	router.POST("/animals", CreateAnimal)
	router.PUT("/animals/:id", UpdateAnimal)
	router.DELETE("/animals/:id", DeleteAnimal)
}
