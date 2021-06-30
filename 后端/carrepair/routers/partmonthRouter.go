package routers

import (
	"carrepair/controllers"
	"github.com/gin-gonic/gin"
)

func partmonthRouter(r *gin.RouterGroup) {
	r.GET("/partmonth", controllers.Getpartmonth)
}

