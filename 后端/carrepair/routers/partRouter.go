package routers

import (
	"carrepair/controllers"
	"github.com/gin-gonic/gin"
)

func partRouter(r *gin.RouterGroup) {
	r.GET("/part", controllers.Getpart)
	r.DELETE("/part", controllers.Delpart)
	r.PUT("/part", controllers.Updatepart)
	r.POST("/part", controllers.Addpart)
}
