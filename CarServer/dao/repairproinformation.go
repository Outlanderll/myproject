package dao

import (
	"carrepair/models"
)
//通过gorm对数据库进行增删改添操作
func Selectrepairpro() ([]models.Repairpro, error){
	var repairpro []models.Repairpro
	err := db.Table("repairpro").Order("proid ASC").Find(&repairpro).Error//查询所有数据并按照升序输出
	return repairpro ,err
}

func Addrepairpro(repairpro models.Repairpro) error {
	err := db.Table("repairpro").Create(&repairpro).Error//插入一行从前端发送过来的数据并且向数据库里添加新的元组
	return err
}
func Updaterepairpro(repairpro models.Repairpro) error{
	err := db.Table("repairpro").Where("proid = ?",repairpro.Proid).Updates(&repairpro).Error//根据前端传来的数据选择与其车牌号一致的元组进行更改
	return err
}
func Deleterepairpro(id string)  error {
	err := db.Table("repairpro").Where("proid = ?",id).Delete(models.Repairpro{}).Error//根据前端输入的id进行删除
	return err
}