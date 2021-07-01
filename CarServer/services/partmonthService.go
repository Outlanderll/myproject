package services

import (
	"carrepair/dao"
	"carrepair/models"
)

func Getpartmonth() (error, []models.Partmonth) {
	partmonthData,err  := dao.Selectpartmonth()
	return err, partmonthData
}
