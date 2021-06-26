package routers

import (
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)

func orderreportRouter(r *gin.RouterGroup) {
	r.GET("/order_report", controllers.Getorderreport)
}
