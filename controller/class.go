package controller

import (
	"github.com/gin-gonic/gin"
)

func GetClass(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get class",
	})
}

func AddClass(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "add class",
	})
}
