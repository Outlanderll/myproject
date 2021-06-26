package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllRepairman() (error, []modules.Repairman) {
	err, RepairmanData := underly.GetAllRepairman()
	return err, RepairmanData
}

func DelRepairman(rno int) error {
	err := underly.DelRepairman(rno)
	return err
}

func UpdateRepairman(repairman modules.Repairman) error {
	err := underly.UpdateRepairman(repairman)
	return err
}

func AddRepairman(repairman modules.Repairman) (error) {
	err := underly.AddRepairman(repairman)
	return err
}