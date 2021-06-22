package underly

import (
	"awesomeProject1/modules"
)

func GetAllRepairman(repairmanParam map[string]interface{}) (error, []modules.Repairman, int64) {
	var repairmanData []modules.Repairman
	page := repairmanParam["page"].(int)
	pageSize := repairmanParam["limit"].(int)
	searchName := repairmanParam["searchName"].(string)
	var total int64
	err := db.Table("repairman").Select("rno, rname, rphonenum, hourly_wage").Where("repairman.rname like ? ", searchName+"%").Order("rno ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&repairmanData).Error
	return err, repairmanData, total
}

func DelRepairman(rno string) error {
	err := db.Table("repairman").Where("rno = ?", rno).Error
	return err
}

func UpdateRepairman(repairman modules.Repairman) error {
	err := db.Table("repairman").Where("rno = ?", repairman.Rno).Updates(&repairman).Error
	return err
}

func AddRepairman(repairman modules.Repairman) (error) {
	err := db.Table("repairman").Select("rno", "rname", "rphonenum", "hourly_wage").Create(&repairman).Error
	return err
}