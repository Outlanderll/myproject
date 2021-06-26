package controllers

import (
	"awesomeProject1/modules"
	"awesomeProject1/pkg/app"
	"awesomeProject1/pkg/e"
	"awesomeProject1/services"
	"encoding/json"
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
	id := -1
	if arg := c.Query("Cno"); arg != " " {
		id = com.StrTo(arg).MustInt()
	}
	if id == -1{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.DelCustomer(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"删除成功")
}

func Updatecustomer(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)
	if m["Cno"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	cno := -1
	cno = com.StrTo(m["Cno"]).MustInt()
	cname := m["Cname"]
	cphonenum := m["Cphonenum"]
	cartype := m["Cartype"]
	licensenumber := m["LicenseNumber"]
	err := services.UpdateCustomer(modules.Customer{Cno: cno, Cname: cname, Cphonenum: cphonenum, Cartype: cartype, LicenseNumber: licensenumber})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addcustomer(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)

	cno := com.StrTo(m["Cno"]).MustInt()
	cname := m["Cname"]
	cphonenum := m["Cphonenum"]
	cartype := m["Cartype"]
	licensenumber := m["LicenseNumber"]
	data := modules.Customer{Cno: cno, Cname: cname, Cphonenum: cphonenum, Cartype: cartype, LicenseNumber: licensenumber}
	err := services.AddCustomer(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"添加成功！")
}