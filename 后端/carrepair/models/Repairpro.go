package models

type Repairpro struct {
	Proid string `gorm:"colunmn :proid" json:"proid"`
	Proname string `gorm:"colunmn :proname" json:"proname"`
	Repairtime int `gorm:"colunmn :repairtine" json:"repairtime"`
	Outnum int `gorm:"colunmn :outnum" json:"outnum"`
	Outtime string `gorm:"colunmn :outtime" json:"outtime"`
	Partid string `gorm:"colunmn :partid" json:"partid"`
	Repairmanid string `gorm:"colunmn :repairmanid" json:"repairmanid"`
	Carid string `gorm:"column :carid" json:"carid"`
}
