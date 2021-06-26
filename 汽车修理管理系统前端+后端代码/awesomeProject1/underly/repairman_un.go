package underly

import (
	"awesomeProject1/modules"
)

func GetAllRepairman() (error, []modules.Repairman) {
	var repairmanData []modules.Repairman
	err := db.Table("repairman").Order("rno ASC").Find(&repairmanData).Error
	return err, repairmanData
}

func DelRepairman(rno int) error {
	err := db.Table("repairman").Where("rno = ?", rno).Delete(modules.Repairman{}).Error
	return err
}

func UpdateRepairman(repairman modules.Repairman) error {
	err := db.Table("repairman").Where("rno = ?", repairman.Rno).Updates(&repairman).Error
	return err
}

func AddRepairman(repairman modules.Repairman) (error) {
	err := db.Table("repairman").Create(&repairman).Error
	return err
}