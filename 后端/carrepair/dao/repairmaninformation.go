package dao

import (
	"carrepair/models"
)
//通过gorm对数据库进行增删改添操作
func Selectrepairman() ([]models.Repairman, error){
	var repairman []models.Repairman
	err := db.Table("repairman").Order("repairmanid ASC").Find(&repairman).Error//查询所有数据并按照升序输出
	return repairman ,err
}

func Addrepairman(repairman models.Repairman) error {
	err := db.Table("repairman").Create(&repairman).Error//插入一行从前端发送过来的数据并且向数据库里添加新的元组
	return err
}
func Updaterepairman(repairman models.Repairman) error{
	err := db.Table("repairman").Where("repairmanid = ?",repairman.Repairmanid).Updates(&repairman).Error//根据前端传来的数据选择与其车牌号一致的元组进行更改
	return err
}
func Deleterepairman(id string)  error {
	err := db.Table("repairman").Where("repairmanid = ?",id).Delete(models.Repairman{}).Error//根据前端输入的id进行删除
	return err
}