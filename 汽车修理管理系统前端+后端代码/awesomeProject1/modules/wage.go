package modules

type Wage struct {
	RepairmanName string `gorm:"column:repairman_name" json:"repairman_name"`
	RepairmanWage int    `gorm:"column:repairman_wage" json:"repairman_wage"`
}
