package services

import (
	"carrepair/dao"
	"carrepair/models"
)

func Getpart() (error, []models.Part) {
	partData,err  := dao.Selectpart()
	return err, partData
}

func Updatepart(part models.Part) error{
	err:= dao.Updatepart(part)
	return err
}

func Addpart(part models.Part) error {
	err := dao.Addpart(part)
	return err
}

func Delpart(id string) error {
	err := dao.Deletepart(id)
	return err
}
