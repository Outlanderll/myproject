package services

import (
	"carrepair/dao"
	"carrepair/models"
)

func Getrepairpro() (error, []models.Repairpro) {
	repairproData,err  := dao.Selectrepairpro()
	return err, repairproData
}

func Updaterepairpro(repairpro models.Repairpro) error{
	err:= dao.Updaterepairpro(repairpro)
	return err
}

func Addrepairpro(repairpro models.Repairpro) error {
	err := dao.Addrepairpro(repairpro)
	return err
}

func Delrepairpro(id string) error {
	err := dao.Deleterepairpro(id)
	return err
}
