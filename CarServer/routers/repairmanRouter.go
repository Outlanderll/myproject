package routers

import (
	"carrepair/controllers"
	"github.com/gin-gonic/gin"
)

func repairmanRouter(r *gin.RouterGroup) {
	r.GET("/repairman", controllers.Getrepairman)
	r.DELETE("/repairman", controllers.Delrepairman)
	r.PUT("/repairman", controllers.Updaterepairman)
	r.POST("/repairman", controllers.Addrepairman)
}