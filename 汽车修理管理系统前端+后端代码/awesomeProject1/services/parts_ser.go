package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllParts() (error, []modules.Parts) {
	err, PartsData := underly.GetAllParts()
	return err, PartsData
}

func DelParts(pno int) error {
	err := underly.DelParts(pno)
	return err
}

func UpdateParts(parts modules.Parts) error {
	err := underly.UpdateParts(parts)
	return err
}

func AddParts(parts modules.Parts) (error) {
	err := underly.AddParts(parts)
	return err
}
