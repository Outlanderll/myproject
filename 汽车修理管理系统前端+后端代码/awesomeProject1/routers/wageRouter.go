package routers

import (
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)

func wageRouter(r *gin.RouterGroup) {
	r.GET("/wage", controllers.Getwage)
}
