package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllWorkingtime() (error, []modules.WorkingTime) {
	err, workingtimeData := underly.GetAllWorkingtime()
	return err, workingtimeData
}
