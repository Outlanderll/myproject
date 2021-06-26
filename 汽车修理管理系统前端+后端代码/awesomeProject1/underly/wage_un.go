package underly

import "awesomeProject1/modules"

func GetAllWage() (error, []modules.Wage) {
	var wageData []modules.Wage
	err := db.Table("wage").Order("repairman_name ASC").Find(&wageData).Error
	return err, wageData
}
