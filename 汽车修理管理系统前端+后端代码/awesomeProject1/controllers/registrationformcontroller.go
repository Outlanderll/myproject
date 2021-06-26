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

func Getregistrationform(c *gin.Context)  {
	err , data := services.GetAllRegistrationform()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delregistrationform(c *gin.Context){
	id := -1
	if arg := c.Query("RegNo"); arg != " " {
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

func Updateregistrationform(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)
	if m["RegNo"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	regno := -1
	regno = com.StrTo(m["RegNo"]).MustInt()
	customerno1 := m["CustomerNo"]
	customerno := com.StrTo(customerno1).MustInt()
	regdate := m["RegDate"]
	regproject := m["RegProject"]
	err := services.UpdateRegistrationform(modules.RegistrationForm{RegNo: regno, CustomerNo: customerno, RegDate: regdate, RegProject: regproject})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addregistrationform(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)

	regno := com.StrTo(m["RegNo"]).MustInt()
	customerno1 := m["CustomerNo"]
	customerno := com.StrTo(customerno1).MustInt()
	regdate := m["RegDate"]
	regproject := m["RegProject"]
	data := modules.RegistrationForm{RegNo: regno, CustomerNo: customerno, RegDate: regdate, RegProject: regproject}
	err := services.AddRegistrationform(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"添加成功！")
}
