package underly

import "awesomeProject1/modules"

func GetAllRepairform() (error, []modules.RepairForm) {
	var repairformData []modules.RepairForm

	err := db.Table("repair_form").Order("rep_no ASC").Find(&repairformData).Error
	return err, repairformData
}

func DelRepairform(reno int) error {
	err := db.Table("repair_form").Where("rep_no = ?", reno).Delete(modules.RepairForm{}).Error
	return err
}

func UpdateRepairform(repairform modules.RepairForm) error {
	err := db.Table("repair_form").Where("rep_no = ?", repairform.RepNo).Updates(&repairform).Error
	return err
}

func AddRepairform(repairform modules.RepairForm) (error) {
	err := db.Table("repair_form").Create(&repairform).Error
	return err
}
