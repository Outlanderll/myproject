package controllers

import (
	"awesomeProject1/modules"
	"awesomeProject1/pkg/app"
	"awesomeProject1/pkg/e"
	"awesomeProject1/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getrepairman(c *gin.Context)  {
	err , data := services.GetAllRepairman()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delrepairman(c *gin.Context){
	id := "-1"
	id = c.PostForm("cno")
	if id == "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	rno := com.StrTo(id).MustInt()
	err := services.DelCustomer(rno)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"删除成功")
}

func Updaterepairman(c *gin.Context){
	rid := c.PostForm("cno")
	rno := com.StrTo(rid).MustInt()
	rname := c.PostForm("cname")
	rphonenum := c.PostForm("rphonenum")
	hourlywage1 := c.PostForm("hourly_wage")
	hourlywage := com.StrTo(hourlywage1).MustInt()
	err := services.UpdateRepairman(modules.Repairman{rno, rname, rphonenum, hourlywage})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addrepairman(c *gin.Context){
	rid := c.PostForm("cno")
	rno := com.StrTo(rid).MustInt()
	rname := c.PostForm("cname")
	rphonenum := c.PostForm("rphonenum")
	hourlywage1 := c.PostForm("hourly_wage")
	hourlywage := com.StrTo(hourlywage1).MustInt()
	a := modules.Repairman{rno, rname, rphonenum, hourlywage}
	err := services.AddRepairman(a)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":a},"添加成功！")
}
