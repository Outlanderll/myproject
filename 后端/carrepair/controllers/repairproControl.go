package controllers

import (
	"carrepair/models"
	"carrepair/pkg/app"
	"carrepair/pkg/e"
	"carrepair/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getrepairpro(c *gin.Context)  {
	err , data := services.Getrepairpro()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delrepairpro(c *gin.Context){
	id := "-1"
	id = c.Query("proid")
	if id == "-1"{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.Delrepairpro(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":id},"删除成功")
}

func Updaterepairpro(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string] string
	_ =json.Unmarshal(a,&m)
	proid := m["proid"]
	proname := m["proname"]
	repairtime := com.StrTo(m["repairtime"]).MustInt()
	outnum := com.StrTo(m["outnum"]).MustInt()
	outtime := m["outtime"]
	partid := m["partid"]
	repairmanid := m["repairmanid"]
	carid := m["carid"]

	err := services.Updaterepairpro(models.Repairpro{proid,proname,repairtime,outnum, outtime, partid,repairmanid,carid})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addrepairpro(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string] string
	_ =json.Unmarshal(a,&m)
	proid := m["proid"]
	proname := m["proname"]
	repairtime := com.StrTo(m["repairtime"]).MustInt()
	outnum := com.StrTo(m["outnum"]).MustInt()
	outtime := m["outtime"]
	partid := m["partid"]
	repairmanid := m["repairmanid"]
	carid := m["carid"]

	err := services.Addrepairpro(models.Repairpro{proid,proname,repairtime,outnum, outtime, partid,repairmanid,carid})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"添加成功！")
}

