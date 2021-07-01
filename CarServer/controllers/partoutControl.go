package controllers

import (
	"carrepair/models"
	"carrepair/pkg/app"
	"carrepair/pkg/e"
	"carrepair/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getpartout(c *gin.Context)  {
	err , data := services.Getpartout()
	if err != nil{
		app.Error(c, e.ERROR,err,err.Error())
	}
	app.OK(c, gin.H{"value": data},"查询成功")
}

func Delpartout(c *gin.Context){
	id := "-1"
	id = c.Query("partid")
	//fmt.Println(id)
	if id == ""{
		app.INFO(c,30001,"参数错误")
		return
	}
	err := services.Delpartout(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{"value":id},"删除成功")
}

func Updatepartout(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string] string
	_ =json.Unmarshal(a,&m)
	partid := m["partid"]
	partname := m["partname"]
	partprice := com.StrTo(m["partprice"]).MustInt()
	outtime :=m["outtime"]
	outnum := com.StrTo(m["outnum"]).MustInt()

	err := services.Updatepartout(models.Partout{partid, partname, partprice, outtime,outnum})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"更新成功！")
}

func Addpartout(c *gin.Context){
	a,_ := c.GetRawData()
	var m map[string] string
	_ =json.Unmarshal(a,&m)
	partid := m["partid"]
	partname := m["partname"]
	partprice := com.StrTo(m["partprice"]).MustInt()
	outtime :=m["outtime"]
	outnum := com.StrTo(m["outnum"]).MustInt()
	fmt.Println(outnum)

	err := services.Addpartout(models.Partout{partid, partname, partprice, outtime,outnum})
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, gin.H{},"添加成功！")
}
