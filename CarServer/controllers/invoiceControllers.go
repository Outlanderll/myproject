package controllers

import (
	"carrepair/pkg/app"
	"carrepair/pkg/e"
	"carrepair/services"
	"github.com/gin-gonic/gin"
)

func Getinvoice(c *gin.Context)  {
	err , data := services.Getinvoice()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}
