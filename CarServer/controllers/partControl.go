package controllers

import (
	"carrepair/models"
	"carrepair/pkg/app"
	"carrepair/pkg/e"
	"carrepair/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getpart(c *gin.Context)  {
	err , data := services.Getpart()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delpart(c *gin.Context){
	data := "-1"
	data = c.Query("partid")
	c.ShouldBindJSON(&data)
	if data== "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.Delpart(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"删除成功")
}

func Updatepart(c *gin.Context){
	var data models.Partreceive
	c.ShouldBindJSON(&data)
	var datanew models.Part
	datanew.Partid = data.Partid
	datanew.Partname = data.Partname
	datanew.Partprice = com.StrTo(data.Partprice).MustInt()
	datanew.Entertime = data.Entertime
	datanew.Enternum = com.StrTo(data.Enternum).MustInt()
	err := services.Updatepart(datanew)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":datanew},"更新成功！")
}

func Addpart(c *gin.Context){
	var data models.Partreceive
	err := c.ShouldBindJSON(&data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	fmt.Println(data)
	var datanew models.Part
	datanew.Partid = data.Partid
	datanew.Partname = data.Partname
	datanew.Partprice = com.StrTo(data.Partprice).MustInt()
	datanew.Entertime = data.Entertime
	datanew.Enternum = com.StrTo(data.Enternum).MustInt()
	err = services.Addpart(datanew)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":datanew},"添加成功！")
}
