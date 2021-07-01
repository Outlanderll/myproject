package models

type Part struct {
	Partid string `gorm:"colunmn :partid" json:"partid"`
	Partname string `gorm:"colunmn :partname" json:"partname"`
	Partprice int `gorm:"colunmn :partprice" json:"partprice"`
	Entertime string `gorm:"colunmn :entertime" json:"entertime"`
	Enternum int `gorm:"colunmn :enternum" json:"enternum"`
}
