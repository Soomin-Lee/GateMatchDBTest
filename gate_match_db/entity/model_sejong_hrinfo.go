package entity

type EmpPhoto struct {
	ID        int          `gorm:"primary_key"`
	Feature   FeatureDBNew `gorm:"ForeignKey:FeatureID;"`
	FeatureID int          `gorm:"not null;"`
	Image     []byte       `gorm:"type:longblob; not null;"`
}

type EmpInfo struct {
	ID           int    `gorm:"primary_key"`
	Emp_no       string `gorm:"not null;"`
	Emp_nm       string `gorm:"not null;"`
	Card_no      string `gorm:"not null;"`
	Comp_cd      string `gorm:"not null;"`
	Comp_nm      string `gorm:"not null;"`
	Dept_cd      string `gorm:"not null;"`
	Dept_nm      string `gorm:"not null;"`
	Posi_nm      string `gorm:"not null;"`
	Regstatus_cd string `gorm:"not null;"`
	Emptype_cd   string `gorm:"not null;"`
}
