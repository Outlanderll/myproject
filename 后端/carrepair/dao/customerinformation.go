package dao

import (
	"carrepair/models"
)
//通过gorm对数据库进行增删改添操作
func Selectcustomer() ([]models.Customer, error){
	var customer []models.Customer
	err := db.Table("customer").Order("carid ASC").Find(&customer).Error//查询所有数据并按照升序输出
	return customer ,err
}

func Addcustomer(customer models.Customer) error {
	err := db.Table("customer").Create(&customer).Error//插入一行从前端发送过来的数据并且向数据库里添加新的元组
	return err
}
func Updatecustomer(customer models.Customer) error{
	err := db.Table("customer").Where("carid = ?",customer.Carid).Updates(&customer).Error//根据前端传来的数据选择与其车牌号一致的元组进行更改
	return err
}
func Deletecustomer(id string)  error {
	err := db.Table("customer").Where("carid = ?",id).Delete(models.Customer{}).Error//根据前端输入的id进行删除
	return err
}