package models

type Partmonth struct {
	Partid string `gorm:"colunmn :partid" json:"partid"`
	Partname string `gorm:"colunmn :partname" json:"partname"`
	Partprice int `gorm:"colunmn :partprice" json:"partprice"`
	Enternum int `gorm:"colunmn :enternum" json:"enternum"`
	Entertime string `gorm:"colunmn :entertime" json:"entertime"`
	Outnum int `gorm:"colunmn :enternum" json:"outnum"`
	Outtime string `gorm:"colunmn :entertime" json:"outtime"`
}
