package controllers

import (
	"awesomeProject1/modules"
	"awesomeProject1/pkg/app"
	"awesomeProject1/pkg/e"
	"awesomeProject1/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getparts(c *gin.Context)  {
	err, data:= services.GetAllParts()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delparts(c *gin.Context){
	id := "-1"
	id = c.PostForm("pno")
	if id == "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	pno := com.StrTo(id).MustInt()
	err := services.DelParts(pno)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"删除成功")
}

func Updateparts(c *gin.Context){
	pid := c.PostForm("pno")
	pno := com.StrTo(pid).MustInt()
	pname := c.PostForm("pname")
	unitprice1 := c.PostForm("unit_price")
	unitprice := com.StrTo(unitprice1).MustInt()
	inventory1 := c.PostForm("inventory")
	inventory := com.StrTo(inventory1).MustInt()
	mininventory1 := c.PostForm("min_inventory")
	mininventory := com.StrTo(mininventory1).MustInt()
	err := services.UpdateParts(modules.Parts{pno, pname, unitprice, inventory, mininventory})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addparts(c *gin.Context){
	pid := c.PostForm("pno")
	pno := com.StrTo(pid).MustInt()
	pname := c.PostForm("pname")
	unitprice1 := c.PostForm("unit_price")
	unitprice := com.StrTo(unitprice1).MustInt()
	inventory1 := c.PostForm("inventory")
	inventory := com.StrTo(inventory1).MustInt()
	mininventory1 := c.PostForm("min_inventory")
	mininventory := com.StrTo(mininventory1).MustInt()
	a := modules.Parts{pno, pname, unitprice, inventory, mininventory}
	err := services.AddParts(a)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":a},"添加成功！")
}
