package modules

type Invoice struct {
	CustomerName  string	`gorm:"column:customer_name" json:"customer_name"`
	LicenseNumber string	`gorm:"column:license_number" json:"license_number"`
	RepProject    string    `gorm:"column:rep_project" json:"rep_project"`
	Cost		  int    	`gorm:"column:cost" json:"cost"`
}