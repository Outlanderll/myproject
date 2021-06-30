package dao

import (
	"carrepair/models"
)
//通过gorm对数据库进行增删改添操作
func Selectinvoice() ([]models.Invoice, error){
	var invoice []models.Invoice
	err := db.Table("invoice").Order("carid ASC").Find(&invoice).Error//查询所有数据并按照升序输出
	return invoice ,err
}
