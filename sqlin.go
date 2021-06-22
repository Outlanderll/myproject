package underly

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Sqlin() {
	var err error
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=carrepair password=138333 sslmode=disable")
	if err != nil {
		fmt.Println("open postgres failed")
		return
	}
	fmt.Println("连接数据库成功!")
}