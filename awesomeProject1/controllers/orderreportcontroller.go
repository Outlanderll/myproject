package controllers

import (
	"awesomeProject1/pkg/app"
	"awesomeProject1/pkg/e"
	"awesomeProject1/services"
	"github.com/gin-gonic/gin"
)

func Getorderreport(c *gin.Context)  {
	err , data := services.GetAllOrderreport()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}