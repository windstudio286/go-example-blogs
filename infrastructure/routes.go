package infrastructure

import (
	docs "blogs/docs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	Gin *gin.Engine
}

func NewGinRouter() GinRouter {
	httpRouter := gin.Default()
	docs.SwaggerInfo.Title = "API"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.BasePath = "/api/v1"
	httpRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Up and Running...!"})
	})

	return GinRouter{
		Gin: httpRouter,
	}
}
