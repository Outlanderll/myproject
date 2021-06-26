package routers

import (
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)

func invoiceRouter(r *gin.RouterGroup) {
	r.GET("/invoice", controllers.Getinvoice)
}
