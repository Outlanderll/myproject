package routers

import (
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)

func workingtimeRouter(r *gin.RouterGroup) {
	r.GET("/working_time", controllers.Getworkingtime)
}
