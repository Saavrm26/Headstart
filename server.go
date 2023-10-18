package main

import (
	"github.com/gin-gonic/gin"
	"hs/headstart/Database"
	"hs/headstart/Routers"
	"log"
)

func main() {
	if err := Database.InitializeFirebase(); err != nil {
		log.Fatal(err)
		return
	}
	initCatalougeValidatorFreeList()
	router := gin.Default()
	Routers.CatalougeRoutes(router)
	router.Run(":8080")
}
