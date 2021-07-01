package services

import (
	"carrepair/dao"
	"carrepair/models"
)

func Getpartout() (error, []models.Partout) {
	partoutData,err  := dao.Selectpartout()
	return err, partoutData
}

func Updatepartout(partout models.Partout) error{
	err:= dao.Updatepartout(partout)
	return err
}

func Addpartout(partout models.Partout) error {
	err := dao.Addpartout(partout)
	return err
}

func Delpartout(id string) error {
	err := dao.Deletepartout(id)
	return err
}
