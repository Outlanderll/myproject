package modules

type OunboundList struct {
	PartsName string `gorm:"column:parts_name" json:"parts_name"`
	Number    int    `gorm:"column:number" json:"number"`
}
