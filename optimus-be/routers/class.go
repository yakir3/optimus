package routers

import (
	"optimus-be/controller"

	"github.com/gin-gonic/gin"
)

func InitClassRoute(r *gin.Engine) {
	classRoute := r.Group("/class")

	classRoute.GET("/", controller.GetClass)
	classRoute.POST("/add", controller.AddClass)
}
