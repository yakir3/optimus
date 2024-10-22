package controller

import (
	"github.com/gin-gonic/gin"
)

func GetStudent(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get student",
	})
}

func AddStudent(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "add student",
	})
}
