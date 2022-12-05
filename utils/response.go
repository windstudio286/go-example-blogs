package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"msg" example:"status bad request"`
}

type HTTPSucess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"msg" example:"status ok"`
}

func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"error": data})
}

func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"msg": data})
}
