package underly

import (
	"awesomeProject1/modules"
)

func GetAllCustomer(customerParam map[string]interface{}) (error, []modules.Customer, int64) {
	var customerData []modules.Customer
	page := customerParam["page"].(int)
	pageSize := customerParam["limit"].(int)
	searchName := customerParam["searchName"].(string)
	var total int64
	err := db.Table("customer c").Select("c.cno, c.cname, c.cphonenum, c.car_type, c.license_number").Joins("INNER JOIN registration r ON c.cno = r.cno").Where("c.cname like ? ", searchName+"%").Order("cno ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&customerData).Error
	return err, customerData, total
}

func DelCustomer(cno string) error {
	err := db.Table("customer").Where("cno = ?", cno).Error
	return err
}

func UpdateCustomer(customer modules.Customer) error {
	err := db.Table("customer").Where("cno = ?", customer.Cno).Updates(&customer).Error
	return err
}

func AddCustomer(customer modules.Customer) (error) {
	err := db.Table("customer").Select("cno", "cname", "cphonenum", "car_type", "license_number").Create(&customer).Error
	return err
}