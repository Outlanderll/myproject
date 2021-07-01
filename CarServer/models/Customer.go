package models

type Customer struct {
	Carid string `gorm:"colunmn :carid" json:"carid"`
	Cusname string `gorm:"colunmn :cusname" json:"cusname"`
	Cuspthone string `gorm:"colunmn :cusphone" json:"cuspthone"`
	Cartype string  `gorm:"colunmn :cartype" json:"cartype"`
}
