package routers

import (
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)

func partRouter(r *gin.RouterGroup) {
	r.GET("/parts", controllers.Getparts)
	r.DELETE("/parts", controllers.Delparts)
	r.PUT("/parts", controllers.Updateparts)
	r.POST("/parts", controllers.Addparts)
}
