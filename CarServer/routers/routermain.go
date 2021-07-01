package routers

import (
	"carrepair/middleWare"
	"github.com/gin-gonic/gin"
)

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/apis")
	customerRouter(r)
	repairmanRouter(r)
	partRouter(r)
	repairproRouter(r)
	partoutRouter(r)
	partinventoryRouter(r)
	partmonthRouter(r)
	invoiceRouter(r)
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleWare.Cors())
	g := r.Group("/carrepair")

	sysNoCheckRoleRouter(g)

	return r
}