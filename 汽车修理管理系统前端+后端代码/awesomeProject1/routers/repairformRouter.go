package routers

import (
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)

func repairformRouter(r *gin.RouterGroup) {
	r.GET("/repair_form", controllers.Getrepairform)
	r.DELETE("/repair_form", controllers.Delrepairform)
	r.PUT("/repair_form", controllers.Updaterepairform)
	r.POST("/repair_form", controllers.Addrepairform)
}
