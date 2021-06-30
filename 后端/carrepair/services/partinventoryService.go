package services

import (
	"carrepair/dao"
	"carrepair/models"
)

func Getpartinventory() (error, []models.Partinventory) {
	partinventoryData,err  := dao.Selectpartinventory()
	return err, partinventoryData
}
