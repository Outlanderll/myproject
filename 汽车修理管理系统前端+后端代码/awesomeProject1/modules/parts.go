package modules

type Parts struct {
	Pno       		int		`gorm:"column:pno" json:"pno"`
	Pname     		string	`gorm:"column:pname" json:"pname"`
	UnitPrice 		int		`gorm:"column:unit_price" json:"unit_price"`
	Inventory 		int		`gorm:"column:inventory" json:"inventory"`
	Mininventory 	int		`gorm:"column:min_inventory" json:"min_inventory"`
}