package modules

type WorkingTime struct {
	RepairmanNo int `gorm:"column:repairman_no" json:"repairman_no"`
	Time        int `gorm:"column:time" json:"time"`
}
