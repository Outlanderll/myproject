package routers

import (
	"carrepair/controllers"
	"github.com/gin-gonic/gin"
)

func customerRouter(r *gin.RouterGroup) {
	r.GET("/customer", controllers.Getcustomer)
	r.DELETE("/customer", controllers.Delcustomer)
	r.PUT("/customer", controllers.Updatecustomer)
	r.POST("/customer", controllers.Addcustomer)
}