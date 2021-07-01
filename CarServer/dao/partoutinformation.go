package dao

import (
	"carrepair/models"
	"errors"
)
//通过gorm对数据库进行增删改添操作
func Selectpartout() ([]models.Partout, error){
	var partout []models.Partout
	err := db.Table("partout").Order("partid ASC").Find(&partout).Error//查询所有数据并按照升序输出
	return partout ,err
}

func Addpartout(partout models.Partout) error {
	var part models.Part
	db.Table("part").Where("partid = ?",partout.Partid).First(&part)
	if part.Partid == ""{
		return errors.New("没有该项零件")
	}
	err := db.Table("partout").Create(&partout).Error//插入一行从前端发送过来的数据并且向数据库里添加新的元组
	if part.Enternum < partout.Outnum{

		return errors.New("出库数量大于已有数量")
	}
	return err
}
func Updatepartout(partout models.Partout) error{
	var part models.Part
	db.Table("part").Where("partid = ?",partout.Partid).First(&part)
	if part.Partid == ""{
		return errors.New("没有该项零件")
	}
	err := db.Table("partout").Where("partid = ?",partout.Partid).Updates(&partout).Error
	if part.Enternum < partout.Outnum{
		return errors.New("出库数量大于已有数量")
	}
	return err
}
func Deletepartout(id string)  error {
	err := db.Table("partout").Where("partid = ?",id).Delete(models.Partout{}).Error//根据前端输入的id进行删除
	return err
}
