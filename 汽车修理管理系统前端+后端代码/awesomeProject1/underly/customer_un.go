package underly

import (
	"awesomeProject1/modules"
)

func GetAllCustomer() (error, []modules.Customer) {
	var customerData []modules.Customer
	err := db.Table("customer").Order("cno ASC").Find(&customerData).Error
	return err, customerData
}

func DelCustomer(cno int) error {
	err := db.Table("customer").Where("cno = ?", cno).Delete(modules.Customer{}).Error
	return err
}

func UpdateCustomer(customer modules.Customer) error {
	err := db.Table("customer").Where("cno = ?", customer.Cno).Updates(&customer).Error
	return err
}

func AddCustomer(customer modules.Customer) error {
	err := db.Table("customer").Create(&customer).Error
	return err
}