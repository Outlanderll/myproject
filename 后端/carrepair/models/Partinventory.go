package models

type Partinventory struct {
	Partid    string `gorm:"colunmn :partid" json:"partid"`
	Partname  string `gorm:"colunmn :partname" json:"partname"`
	Partprice int    `gorm:"colunmn :partprice" json:"partprice"`
	Inventory int    `gorm:"colunmn :inventory" json:"inventory"`
}
