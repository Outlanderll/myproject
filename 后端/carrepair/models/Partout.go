package models

type Partout struct {
	Partid string `gorm:"colunmn :partid" json:"partid"`
	Partname string `gorm:"colunmn :partname" json:"partname"`
	Partprice int `gorm:"colunmn :partprice" json:"partprice"`
	Outtime string `gorm:"colunmn :entertime" json:"outtime"`
	Outnum int `gorm:"colunmn :enternum" json:"outnum"`
}
