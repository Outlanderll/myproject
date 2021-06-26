package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllOrderreport() (error, []modules.OrderReport) {
	err, orderreportData := underly.GetAllOrderreport()
	return err, orderreportData
}

