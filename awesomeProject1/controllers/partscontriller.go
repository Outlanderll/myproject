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

func Getparts(c *gin.Context)  {
	err, data:= services.GetAllParts()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delparts(c *gin.Context){
	id := -1
	if arg := c.Query("Pno"); arg != " " {
		id = com.StrTo(arg).MustInt()
	}
	if id == -1{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.DelParts(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"删除成功")
}

func Updateparts(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)
	if m["Pno"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	pno := -1
	pno = com.StrTo(m["Pno"]).MustInt()
	pname := m["Pname"]
	unitprice1 := m["UnitPrice"]
	unitprice := com.StrTo(unitprice1).MustInt()
	inventory1 := m["Inventory"]
	inventory := com.StrTo(inventory1).MustInt()
	mininventory1 := m["Mininventory "]
	mininventory := com.StrTo(mininventory1).MustInt()
	err := services.UpdateParts(modules.Parts{Pno: pno, Pname: pname, UnitPrice: unitprice, Inventory: inventory, Mininventory: mininventory})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addparts(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)
	pid := m["Pno"]
	pno := com.StrTo(pid).MustInt()
	pname := m["Pname"]
	unitprice1 := m["UnitPrice "]
	unitprice := com.StrTo(unitprice1).MustInt()
	inventory1 := m["Inventory"]
	inventory := com.StrTo(inventory1).MustInt()
	mininventory1 := m["Mininventory "]
	mininventory := com.StrTo(mininventory1).MustInt()
	data := modules.Parts{Pno: pno, Pname: pname, UnitPrice: unitprice, Inventory: inventory, Mininventory: mininventory}
	err := services.AddParts(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"添加成功！")
}
