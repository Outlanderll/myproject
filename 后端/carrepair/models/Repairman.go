package models

type Repairman struct {
	Repairmanid string `gorm:"colunmn :repairmanid"`
	Repairmanname string `gorm:"colunmn :repairmanname"`
	Hourgalary int `gorm:"colunmn :hourgalary"`
}
