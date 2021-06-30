package dao

import (
	"carrepair/models"
)
//通过gorm对数据库进行增删改添操作
func Selectpartmonth() ([]models.Partmonth, error){
	var partmonth []models.Partmonth
	err := db.Table("partmonth").Order("partid ASC").Find(&partmonth).Error//查询所有数据并按照升序输出
	return partmonth ,err
}
