package routers

import (
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)

func outboundlistRouter(r *gin.RouterGroup) {
	r.GET("/outbound_list", controllers.Getoutboundlist)
}
