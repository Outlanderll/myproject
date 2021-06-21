package modules

import "github.com/jinzhu/gorm"

func Delect(db *gorm.DB)  {
	err := db.Table("customer").Where("cname = ?", "马文").Delete("customer").Error
	if err != nil {
		panic(err)
		return
	}
}