package main

import (
	"github.com/gin-gonic/gin"
	"hs/headstart/Routers"
)

func main() {
	initializeFirebase()
	initCatalougeValidatorFreeList()
	router := gin.Default()
	Routers.CatalougeRoutes(router)
	router.Run(":8080")
}
