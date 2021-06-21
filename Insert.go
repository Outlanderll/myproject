package modules

import (
	"github.com/jinzhu/gorm"
)
func Insert(db *gorm.DB)  {
	cus := Customer{Cno:"06", Cname:"马文", Phonenum: "16076698765"}
	db.Table("customer").NewRecord(cus)
	err := db.Table("customer").Create(&cus).Error
	db.Table("customer").NewRecord(cus)
	if err != nil {
		panic(err)
		return
	}
}
