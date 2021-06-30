package routers

import (
	"carrepair/controllers"
	"github.com/gin-gonic/gin"
)

func partoutRouter(r *gin.RouterGroup) {
	r.GET("/partout", controllers.Getpartout)
	r.DELETE("/partout", controllers.Delpartout)
	r.PUT("/partout", controllers.Updatepartout)
	r.POST("/partout", controllers.Addpartout)
}
