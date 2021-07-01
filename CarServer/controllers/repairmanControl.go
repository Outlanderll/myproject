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

func Getrepairman(c *gin.Context)  {
	err , data := services.Getrepairman()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delrepairman(c *gin.Context){

	data := c.Query("repairmanid")
	fmt.Println(data)
	if data == "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.Delrepairman(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"删除成功")
}

func Updaterepairman(c *gin.Context){
	var data models.Repairmanreceive
	var datanew models.Repairman
	c.ShouldBindJSON(&data)
	 datanew.Hourgalary = com.StrTo(data.Hourgalary).MustInt()
	 datanew.Repairmanid = data.Repairmanid
	 datanew.Repairmanname = data.Repairmanname
 	err := services.Updaterepairman(datanew)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":datanew},"更新成功！")
}

func Addrepairman(c *gin.Context){
	var data models.Repairmanreceive
	var datanew models.Repairman
	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}
	datanew.Hourgalary = com.StrTo(data.Hourgalary).MustInt()
	datanew.Repairmanid = data.Repairmanid
	datanew.Repairmanname = data.Repairmanname
	err = services.Addrepairman(datanew)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"添加成功！")
}
