package Controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllCatalouges(ctx *gin.Context) {
  // TODO: fetch info from firestore
	data := make([]map[string]any, 0)
	data = append(data, gin.H{
		"downloadURL": "abc.com",
		"imageURL": "abc.image.com",
		"name": "material theme",
		"description": "next js template",
	})
	data = append(data, gin.H{
		"downloadURL": "abc.com",
		"imageURL": "abc.image.com",
		"name": "material theme",
		"description": "next js template",
	})
	ctx.JSON(http.StatusOK, data)
}


