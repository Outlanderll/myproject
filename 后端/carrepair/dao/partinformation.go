package dao

import (
	"carrepair/models"
)
//通过gorm对数据库进行增删改添操作
func Selectpart() ([]models.Part, error){
	var part []models.Part
	err := db.Table("part").Order("partid ASC").Find(&part).Error//查询所有数据并按照升序输出
	return part ,err
}

func Addpart(part models.Part) error {
	err := db.Table("part").Create(&part).Error//插入一行从前端发送过来的数据并且向数据库里添加新的元组
	return err
}
func Updatepart(part models.Part) error{
	err := db.Table("part").Where("partid = ?",part.Partid).Updates(&part).Error//根据前端传来的数据选择与其车牌号一致的元组进行更改
	return err
}
func Deletepart(id string)  error {
	err := db.Table("part").Where("partid = ?",id).Delete(models.Part{}).Error//根据前端输入的id进行删除
	return err
}