package modules

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Query(db *gorm.DB)  {
	u1 := Customer{}
	err :=	db.Table("customer").Find(&u1).Error
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(u1)
	//err := db.Table("customer").Where("cno='05'").Find(&u1).Error
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
}