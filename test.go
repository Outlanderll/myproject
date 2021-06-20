package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type Customer struct {
	Cno string
	Cname string
	Phonenum string
}

var db *gorm.DB

func insql() {
	database, err := gorm.Open("postgres", "port=5432 user=postgres dbname=carrepair password=138333 sslmode=disable")
	if err != nil {
		fmt.Println("open postgres failed")
		return
	}
	fmt.Println("连接数据库成功!")
	db = database
	//defer db.Close()
}
func query() {

}

func main() {
	insql()
	var u1 Customer
	var u2 []Customer
	db.Table("customer").Find(&u2)
	err := db.Table("customer").Where("cno='02'").Find(&u1).Error
	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("%#v",u1)
	fmt.Println(u2)
	//defer db.Close()
}