package modules

type RegistrationForm struct {
	RegNo      int		`gorm:"column:reg_no" json:"reg_no"`
	CustomerNo int 		`gorm:"column:customer_no" json:"customer_no"`
	RegDate    string 	`gorm:"column:reg_date" json:"reg_date"`
	RegProject string 	`gorm:"column:reg_progect" json:"reg_progect"`
}
