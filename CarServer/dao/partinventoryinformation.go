package dao

import (
	"carrepair/models"
)
//通过gorm对数据库进行增删改添操作
func Selectpartinventory() ([]models.Partinventory, error){
	var partinventory []models.Partinventory
	err := db.Table("partinventory").Order("partid ASC").Find(&partinventory).Error//查询所有数据并按照升序输出
	return partinventory ,err
}
