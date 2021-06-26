package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllRegistrationform() (error, []modules.RegistrationForm) {
	err, registrationformData := underly.GetAllRegistrationform()
	return err, registrationformData
}

func DelRegistrationform(regno int) error {
	err := underly.DelRegistrationform(regno)
	return err
}

func UpdateRegistrationform(registrationform modules.RegistrationForm) error {
	err := underly.UpdateRegistrationform(registrationform)
	return err
}

func AddRegistrationform(registrationform modules.RegistrationForm) (error) {
	err := underly.AddRegistrationform(registrationform)
	return err
}