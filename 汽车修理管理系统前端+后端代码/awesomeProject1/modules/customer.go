package modules

type Customer struct {
	Cno           int		`gorm:"column:cno" json:"cno" `
	Cname         string	`gorm:"column:cname" json:"cname"`
	Cphonenum     string	`gorm:"column:cphonenum" json:"cphonenum"`
	Cartype       string	`gorm:"column:car_type" json:"car_type"`
	LicenseNumber string	`gorm:"column:license_number" json:"license_number"`
}

