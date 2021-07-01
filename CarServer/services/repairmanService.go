package services

import (
	"carrepair/dao"
	"carrepair/models"
)

func Getrepairman() (error, []models.Repairman) {
	repairmanData,err  := dao.Selectrepairman()
	return err, repairmanData
}

func Updaterepairman(repairman models.Repairman) error{
	err:= dao.Updaterepairman(repairman)
	return err
}

func Addrepairman(repairman models.Repairman) error {
	err := dao.Addrepairman(repairman)
	return err
}

func Delrepairman(id string) error {
	err := dao.Deleterepairman(id)
	return err
}