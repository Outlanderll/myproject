package underly

import "awesomeProject1/modules"

func GetAllParts() (error, []modules.Parts) {
	var partsData []modules.Parts

	err := db.Table("parts").Order("pno ASC").Find(&partsData).Error
	return err, partsData
}

func DelParts(pno int) error {
	err := db.Table("parts").Where("pno = ?", pno).Delete(modules.Parts{}).Error
	return err
}

func UpdateParts(parts modules.Parts) error {
	err := db.Table("parts").Where("pno = ?", parts.Pno).Updates(&parts).Error
	return err
}

func AddParts(parts modules.Parts) error {
	err := db.Table("parts").Create(&parts).Error
	return err
}
