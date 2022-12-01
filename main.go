package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello world !"})
	})
	router.Run(":8000")
}
