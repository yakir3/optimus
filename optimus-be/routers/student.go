package routers

import (
	"optimus-be/controller"

	"github.com/gin-gonic/gin"
)

func InitStudentRoute(r *gin.Engine) {
	stuRoute := r.Group("/student")

	stuRoute.GET("/", controller.GetStudent)
	stuRoute.POST("/add", controller.AddStudent)
}
