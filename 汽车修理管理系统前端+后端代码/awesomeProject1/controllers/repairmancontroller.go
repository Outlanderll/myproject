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

func Getrepairman(c *gin.Context)  {
	err , data := services.GetAllRepairman()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delrepairman(c *gin.Context){
	id := -1
	if arg := c.Query("Rno"); arg != " " {
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

func Updaterepairman(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)
	if m["Rno"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	rno := -1
	rno = com.StrTo("Rno").MustInt()
	rname := m["Rname"]
	rphonenum := m["Rphonenum"]
	hourlywage1 := m["HourlyWage"]
	hourlywage := com.StrTo(hourlywage1).MustInt()
	err := services.UpdateRepairman(modules.Repairman{Rno: rno, Rname: rname, Rphonenum: rphonenum, HourlyWage: hourlywage})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addrepairman(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)
	rno := com.StrTo("Rno").MustInt()
	rname := m["Rname"]
	rphonenum := m["Rphonenum"]
	hourlywage1 := m["HourlyWage"]
	hourlywage := com.StrTo(hourlywage1).MustInt()
	data := modules.Repairman{Rno: rno, Rname: rname, Rphonenum: rphonenum, HourlyWage: hourlywage}
	err := services.AddRepairman(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"添加成功！")
}
