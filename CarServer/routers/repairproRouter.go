package routers

import (
	"carrepair/controllers"
	"github.com/gin-gonic/gin"
)

func repairproRouter(r *gin.RouterGroup) {
	r.GET("/repairpro", controllers.Getrepairpro)
	r.DELETE("/repairpro", controllers.Delrepairpro)
	r.PUT("/repairpro", controllers.Updaterepairpro)
	r.POST("/repairpro", controllers.Addrepairpro)
}
