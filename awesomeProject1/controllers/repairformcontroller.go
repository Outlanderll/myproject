package controllers

import (
	"awesomeProject1/modules"
	"awesomeProject1/pkg/app"
	"awesomeProject1/pkg/e"
	"awesomeProject1/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getrepairform(c *gin.Context)  {
	err , data := services.GetAllRepairform()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delrepairform(c *gin.Context){
	id := "-1"
	id = c.PostForm("rep_no")
	if id == "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	repno := com.StrTo(id).MustInt()
	err := services.DelCustomer(repno)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"删除成功")
}

func Updaterepairform(c *gin.Context){
	repid := c.PostForm("rep_no")
	repno := com.StrTo(repid).MustInt()
	customerno1 := c.PostForm("customer_no")
	customerno := com.StrTo(customerno1).MustInt()
	repairmanno1 := c.PostForm("repairman_no")
	repairmanno := com.StrTo(repairmanno1).MustInt()
	repproject := c.PostForm("rep_project")
	partsno1 := c.PostForm("parts_no")
	partsno := com.StrTo(partsno1).MustInt()
	partsconsumed1 := c.PostForm("parts_consumed")
	partsconsumed := com.StrTo(partsconsumed1).MustInt()
	repairduration1:= c.PostForm("repair_duration")
 	repairduration := com.StrTo(repairduration1).MustInt()
	err := services.UpdateRepairform(modules.RepairForm{repno, customerno, repairmanno, repproject, partsno, partsconsumed, repairduration})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addrepairform(c *gin.Context){
	repid := c.PostForm("rep_no")
	repno := com.StrTo(repid).MustInt()
	customerno1 := c.PostForm("customer_no")
	customerno := com.StrTo(customerno1).MustInt()
	repairmanno1 := c.PostForm("repairman_no")
	repairmanno := com.StrTo(repairmanno1).MustInt()
	repproject := c.PostForm("rep_project")
	partsno1 := c.PostForm("parts_no")
	partsno := com.StrTo(partsno1).MustInt()
	partsconsumed1 := c.PostForm("parts_consumed")
	partsconsumed := com.StrTo(partsconsumed1).MustInt()
	repairduration1:= c.PostForm("repair_duration")
	repairduration := com.StrTo(repairduration1).MustInt()
	a := modules.RepairForm{repno, customerno, repairmanno, repproject, partsno, partsconsumed, repairduration}
	err := services.AddRepairform(a)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":a},"添加成功！")
}