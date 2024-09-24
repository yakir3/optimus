package route

import (
	"go-pic-tools/controller"

	"github.com/gin-gonic/gin"
)

func InitStudentRoute(r *gin.Engine) {
	stuRoute := r.Group("/student")

	stuRoute.GET("/", controller.GetStudent)
	stuRoute.POST("/add", controller.AddStudent)
}
