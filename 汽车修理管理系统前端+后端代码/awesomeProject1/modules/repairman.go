package modules

type Repairman struct {
	Rno        int		`gorm:"column:rno" json:"rno"`
	Rname      string	`gorm:"column:rname" json:"rname"`
	Rphonenum  string	`gorm:"column:rphonenum" json:"rphonenum"`
	HourlyWage int		`gorm:"column:hourly_wage" json:"hourly_wage"`
}