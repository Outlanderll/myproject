package services

import (
	"carrepair/dao"
	"carrepair/models"
)

func Getcustomer() (error, []models.Customer) {
	customerData,err  := dao.Selectcustomer()
	return err, customerData
}

func Updatecustomer(customer models.Customer) error{
	err:= dao.Updatecustomer(customer)
	return err
}

func Addcustomer(customer models.Customer) error {
	err := dao.Addcustomer(customer)
	return err
}

func Delcustomer(id string) error {
	err := dao.Deletecustomer(id)
	return err
}