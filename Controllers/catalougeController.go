package Controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllCatalouges(ctx *gin.Context) {
  // TODO: fetch info from firestore
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}


