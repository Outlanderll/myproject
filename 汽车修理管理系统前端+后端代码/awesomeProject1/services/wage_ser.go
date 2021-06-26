package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllWage() (error, []modules.Wage) {
	err, wageData := underly.GetAllWage()
	return err, wageData
}
