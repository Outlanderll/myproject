package modules

type RepairForm struct {
	RepNo          int		`gorm:"column:rep_no" json:"rep_no"`
	CustomerNo     int	 	`gorm:"column:customer_no" json:"customer_nocost"`
	RepairmanNo    int	 	`gorm:"column:repairman_no" json:"repairman_no"`
	RepProject     string 	`gorm:"column:rep_project" json:"rep_project"`
	PartsNo        int	 	`gorm:"column:parts_no" json:"parts_no"`
	PartsConsumed  int    	`gorm:"column:parts_consumed" json:"parts_consumed"`
	RepairDuration int    	`gorm:"column:repair_duration" json:"repair_duration"`
}
