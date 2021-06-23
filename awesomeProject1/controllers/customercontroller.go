package controllers

import (
	"awesomeProject1/modules"
	"awesomeProject1/pkg/app"
	"awesomeProject1/pkg/e"
	"awesomeProject1/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getcustomer(c *gin.Context)  {
	err , data := services.GetAllCustomer()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delcustomer(c *gin.Context){
	id := "-1"
	id = c.PostForm("cno")
	if id == "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	cno := com.StrTo(id).MustInt()
	err := services.DelCustomer(cno)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"删除成功")
}

func Updatecustomer(c *gin.Context){
	cid := c.PostForm("cno")
	cno := com.StrTo(cid).MustInt()
	cname := c.PostForm("cname")
	cphonenum := c.PostForm("cphonenum")
	cartype := c.PostForm("car_type")
	licensenumber := c.PostForm("license_number")
	err := services.UpdateCustomer(modules.Customer{cno, cname, cphonenum, cartype, licensenumber})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addcustomer(c *gin.Context){
	cid := c.PostForm("cno")
	cno := com.StrTo(cid).MustInt()
	cname := c.PostForm("cname")
	cphonenum := c.PostForm("cphonenum")
	cartype := c.PostForm("car_type")
	licensenumber := c.PostForm("license_number")
	a := modules.Customer{cno, cname, cphonenum, cartype, licensenumber}
	err := services.AddCustomer(a)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":a},"添加成功！")
}