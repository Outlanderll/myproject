package underly

import "awesomeProject1/modules"

func GetAllWorkingtime() (error, []modules.WorkingTime) {
	var workingtimeData []modules.WorkingTime
	err := db.Table("working_time").Order("repairman_no ASC").Find(&workingtimeData).Error
	return err, workingtimeData
}
