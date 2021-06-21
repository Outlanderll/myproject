package modules

type Customer struct {
	Cno string
	Cname string
	Phonenum string
}
type Vehicle struct {
	Vno     string
	Type    string
	LicNum  string
	DamType string
}
type Repairman struct {
	Rno        string
	Rname      string
	HourlyWage int
	TotalHours int
}
type Parts struct {
	Pno       string
	Pname     string
	UnitPrice int
	Inventory int
}
type Supplier struct {
	Sname string
}

type Sp struct {
	Pno        string
	Sname      string
	SpPrice    int
	SpQuantity int
	SpDate     string
}
type Cv struct {
	Cno    string
	Vno    string
	CvDate string
}
type Vrp struct {
	Vno            string
	Rno            string
	Pno            string
	RepairDate     string
	RepairDuration int
	RepairProject  string
	PartsConsumed  int
	RepairPrice    int
}
