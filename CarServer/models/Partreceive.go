package models
type Partreceive struct {
	Partid string `gorm:"colunmn :partid" json:"partid"`
	Partname string `gorm:"colunmn :partname" json:"partname"`
	Partprice string `gorm:"colunmn :partprice" json:"partprice"`
	Entertime string `gorm:"colunmn :entertime" json:"entertime"`
	Enternum string `gorm:"colunmn :enternum" json:"enternum"`
}