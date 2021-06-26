package underly

import "awesomeProject1/modules"

func GetAllRegistrationform ()(error, []modules.RegistrationForm) {
	var registrationformData []modules.RegistrationForm
	err := db.Table("registration_form").Order("reg_no ASC").Find(&registrationformData).Error
	return err, registrationformData
}

func DelRegistrationform(regno int) error {
	err := db.Table("registration_form").Where("reg_no = ?", regno).Delete(modules.RegistrationForm{}).Error
	return err
}

func UpdateRegistrationform(registrationform modules.RegistrationForm) error {
	err := db.Table("registration_form").Where("reg_no = ?", registrationform.RegNo).Updates(&registrationform).Error
	return err
}

func AddRegistrationform(registrationform modules.RegistrationForm) error {
	err := db.Table("registration_form").Create(&registrationform).Error
	return err
}
