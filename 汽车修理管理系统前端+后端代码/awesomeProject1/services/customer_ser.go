package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllCustomer() (error, []modules.Customer) {
	err, customerData := underly.GetAllCustomer()
	return err, customerData
}

func DelCustomer(cno int) error {
	err := underly.DelCustomer(cno)
	return err
}

func UpdateCustomer(customer modules.Customer) error {
	err := underly.UpdateCustomer(customer)
	return err
}

func AddCustomer(customer modules.Customer) (error) {
	err := underly.AddCustomer(customer)
	return err
}