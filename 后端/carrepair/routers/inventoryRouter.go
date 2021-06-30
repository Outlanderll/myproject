package routers

import (
	"carrepair/controllers"
	"github.com/gin-gonic/gin"
)

func partinventoryRouter(r *gin.RouterGroup) {
	r.GET("/partinventory", controllers.Getpartinventory)
}
