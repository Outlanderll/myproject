package underly

import "awesomeProject1/modules"

func GetAllOrderreport() (error, []modules.OrderReport) {
	var orderreportData []modules.OrderReport
	err := db.Table("order_report").Order("parts_name ASC").Find(&orderreportData).Error
	return err, orderreportData
}
