package routers

import (
	"awesomeProject1/middleWare"
	"github.com/gin-gonic/gin"
)

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/apis")
	customerRouter(r)
	repairmanRouter(r)
	partRouter(r)
	registrationformRouter(r)
	repairformRouter(r)
	invoiceRouter(r)
	orderreportRouter(r)
	outboundlistRouter(r)
	wageRouter(r)
	workingtimeRouter(r)
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleWare.Cors())
	g := r.Group("/system")

	sysNoCheckRoleRouter(g)

	return r
}
