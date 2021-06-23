package controllers

import (
	"awesomeProject1/modules"
	"awesomeProject1/pkg/app"
	"awesomeProject1/pkg/e"
	"awesomeProject1/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getregistrationform(c *gin.Context)  {
	err , data := services.GetAllRegistrationform()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delregistrationform(c *gin.Context){
	id := "-1"
	id = c.PostForm("reg_no")
	if id == "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	regno := com.StrTo(id).MustInt()
	err := services.DelCustomer(regno)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"删除成功")
}

func Updateregistrationform(c *gin.Context){
	regid := c.PostForm("reg_no")
	regno := com.StrTo(regid).MustInt()
	customerno1 := c.PostForm("customer_no")
	customerno := com.StrTo(customerno1).MustInt()
	regdate := c.PostForm("reg_date")
	regproject := c.PostForm("reg_progect")
	err := services.UpdateRegistrationform(modules.RegistrationForm{regno, customerno, regdate, regproject})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addregistrationform(c *gin.Context){
	regid := c.PostForm("reg_no")
	regno := com.StrTo(regid).MustInt()
	customerno1 := c.PostForm("customer_no")
	customerno := com.StrTo(customerno1).MustInt()
	regdate := c.PostForm("reg_date")
	regproject := c.PostForm("reg_progect")
	a := modules.RegistrationForm{regno, customerno, regdate, regproject}
	err := services.AddRegistrationform(a)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":a},"添加成功！")
}
