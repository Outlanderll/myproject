package underly

import "awesomeProject1/modules"

func GetAllRepairform(repairformParam map[string]interface{}) (error, []modules.RepairForm, int64) {
	var repairformData []modules.RepairForm
	page := repairformParam["page"].(int)
	pageSize := repairformParam["limit"].(int)
	searchList := repairformParam["searchName"].(int)
	var total int64
	err := db.Table("repair_form r").Select("r.rep_no, r.cno, r.rno, r.rep_project, r.pno, r.parts_consumed, r.repair_duration").Where("r.rep_no = ? ", searchList).Order("cno ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&repairformData).Error
	return err, repairformData, total
}

func DelRepairform(reno string) error {
	err := db.Table("repair_form").Where("rep_no = ?", reno).Error
	return err
}

func UpdateRepairform(repairform modules.RepairForm) error {
	err := db.Table("repair_form").Where("rep_no = ?", repairform.RepNo).Updates(&repairform).Error
	return err
}

func AddRepairform(repairform modules.RepairForm) (error) {
	err := db.Table("repair_form").Select("rep_no", "cno", "rno", "rep_project", "pno", "parts_consumed", "repair_duration").Create(&repairform).Error
	return err
}
