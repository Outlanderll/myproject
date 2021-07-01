package services

import (
	"carrepair/dao"
	"carrepair/models"
)

func Getinvoice() (error, []models.Invoice) {
	invoiceData,err  := dao.Selectinvoice()
	return err, invoiceData
}
