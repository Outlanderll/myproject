package controllers

import (
	"carrepair/models"
	"carrepair/pkg/app"
	"carrepair/pkg/e"
	"carrepair/services"
	"github.com/gin-gonic/gin"
)

func Getcustomer(c *gin.Context)  {
	err , data := services.Getcustomer()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delcustomer(c *gin.Context){
	data := c.Query("carid")
	if data == "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.Delcustomer(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"删除成功")
}

func Updatecustomer(c *gin.Context){
	var  data models.Customer
	err := c.ShouldBindJSON(&data)
	if err != nil {
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	err = services.Updatecustomer(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"更新成功！")
}

func Addcustomer(c *gin.Context){
	var  data models.Customer
	err := c.ShouldBindJSON(&data)
	if err != nil {
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	err = services.Addcustomer(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"添加成功！")
}
