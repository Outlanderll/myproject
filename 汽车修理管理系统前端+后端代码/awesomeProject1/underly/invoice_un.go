package underly

import "awesomeProject1/modules"

func GetAllInvoice() (error, []modules.Invoice) {
	var invoiceData []modules.Invoice
	err := db.Table("invoice").Order("customer_name ASC").Find(&invoiceData).Error
	return err, invoiceData
}
