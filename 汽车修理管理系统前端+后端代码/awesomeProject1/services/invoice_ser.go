package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllInvoice() (error, []modules.Invoice) {
	err, invoiceData := underly.GetAllInvoice()
	return err, invoiceData
}
