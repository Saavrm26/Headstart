package Routers

import (
	"github.com/gin-gonic/gin"
  "hs/headstart/Controllers"
)

func CatalougeRoutes(router *gin.Engine) {
	catalouge := router.Group("/catalouge")
  catalouge.POST("/", Controllers.PostACatalouge)
	catalouge.GET("/", Controllers.GetAllCatalouges)
}
