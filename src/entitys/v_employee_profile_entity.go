package entitys

type V_employee_profile struct {
	EmpId          string `gorm:"column:emp_id" json:"emp_id"`
	EmpName        string `gorm:"column:emp_name" json:"emp_name"`
	DepartmentName string `gorm:"column:department_name" json:"department_name"`
	UnitName       string `gorm:"column:unit_name" json:"unit_name"`
	AreaName       string `gorm:"column:area_name" json:"area_name"`
	SectorName     string `gorm:"column:sector_name" json:"sector_name"`
	LineName       string `gorm:"column:line_name" json:"line_name"`
	GroupName      string `gorm:"column:group_name" json:"group_name"`
	PositionName   string `gorm:"column:position_name" json:"position_name"`
}

func (V_employee_profile) TableName() string {
	return "v_employee_profile"
}
