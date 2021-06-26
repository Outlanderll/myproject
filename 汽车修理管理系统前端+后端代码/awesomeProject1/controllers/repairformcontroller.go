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

func Getrepairform(c *gin.Context)  {
	err , data := services.GetAllRepairform()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delrepairform(c *gin.Context){
	id := -1
	if arg := c.Query("RepNo"); arg != " " {
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

func Updaterepairform(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)
	if m["RepNo"] == ""{
		app.INFO(c, 30000, "参数非法")
		return
	}
	repno := -1
	repno = com.StrTo(m["RepNo"]).MustInt()
	customerno1 := m["CustomerNo"]
	customerno := com.StrTo(customerno1).MustInt()
	repairmanno1 := m["RepairmanNo"]
	repairmanno := com.StrTo(repairmanno1).MustInt()
	repproject := m["RepProject"]
	partsno1 := m["PartsNo"]
	partsno := com.StrTo(partsno1).MustInt()
	partsconsumed1 := m["PartsConsumed"]
	partsconsumed := com.StrTo(partsconsumed1).MustInt()
	repairduration1:= m["RepairDuration"]
 	repairduration := com.StrTo(repairduration1).MustInt()
	err := services.UpdateRepairform(modules.RepairForm{RepNo: repno, CustomerNo: customerno, RepairmanNo: repairmanno, RepProject: repproject, PartsNo: partsno, PartsConsumed: partsconsumed, RepairDuration: repairduration})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addrepairform(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string]string
	_ =json.Unmarshal(a,&m)
	repno := -1
	repno = com.StrTo(m["RepNo"]).MustInt()
	customerno1 := m["CustomerNo"]
	customerno := com.StrTo(customerno1).MustInt()
	repairmanno1 := m["RepairmanNo"]
	repairmanno := com.StrTo(repairmanno1).MustInt()
	repproject := m["RepProject"]
	partsno1 := m["PartsNo"]
	partsno := com.StrTo(partsno1).MustInt()
	partsconsumed1 := m["PartsConsumed"]
	partsconsumed := com.StrTo(partsconsumed1).MustInt()
	repairduration1:= m["RepairDuration"]
	repairduration := com.StrTo(repairduration1).MustInt()
	data := modules.RepairForm{RepNo: repno, CustomerNo: customerno, RepairmanNo: repairmanno, RepProject: repproject, PartsNo: partsno, PartsConsumed: partsconsumed, RepairDuration: repairduration}
	err := services.AddRepairform(data)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":data},"添加成功！")
}