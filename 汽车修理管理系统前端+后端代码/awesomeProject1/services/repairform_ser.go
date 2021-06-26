package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllRepairform() (error, []modules.RepairForm) {
	err, repairformData := underly.GetAllRepairform()
	return err, repairformData
}

func DelRepairform(repno int) error {
	err := underly.DelRepairform(repno)
	return err
}

func UpdateRepairform(repairform modules.RepairForm) error {
	err := underly.UpdateRepairform(repairform)
	return err
}

func AddRepairform(repairform modules.RepairForm) (error) {
	err := underly.AddRepairform(repairform)
	return err
}