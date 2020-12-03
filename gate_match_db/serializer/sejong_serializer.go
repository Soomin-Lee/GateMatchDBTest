package serializer

type SejongFeatureDBNew struct {
	Emp_no        string
	FeatureVector [512]float32
}

type SejongEmpInfo struct {
	Emp_no       string
	Emp_nm       string
	Card_no      string
	Comp_cd      string
	Comp_nm      string
	Dept_cd      string
	Dept_nm      string
	Posi_nm      string
	Regstatus_cd string
	Emptype_cd   string
}
