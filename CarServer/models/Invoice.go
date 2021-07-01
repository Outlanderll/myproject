package models

type Invoice struct {
	Cusname    string `gorm:"colunmn :cusname" json:"cusname"`
	Carid      string `gorm:"colunmn :carid" json:"carid"`
	Proname    string `grom:"column :proname" json:"proname"`
	Repaircost int    `gorm:"column :repaircost" json:"repaircost"`
}

