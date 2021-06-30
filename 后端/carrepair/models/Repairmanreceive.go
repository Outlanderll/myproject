package models

type Repairmanreceive struct {
	Repairmanid string `gorm:"colunmn :repairmanid" json:"repairmanid"`
	Repairmanname string `gorm:"colunmn :repairmanname" json:"repairmanname"`
	Hourgalary string `gorm:"colunmn :hourgalary" json:"hourgalary"`
}

