package routers

import (
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)

func registrationformRouter(r *gin.RouterGroup) {
	r.GET("/registration_form", controllers.Getregistrationform)
	r.DELETE("/registration_form", controllers.Delregistrationform)
	r.PUT("/registration_form", controllers.Updateregistrationform)
	r.POST("/registration_form", controllers.Addregistrationform)
}
