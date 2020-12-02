package entity

type FeatureDB struct {
	ID          int    `gorm:"primary_key"`
	Emp_no      string `gorm:"not null;"`
	FeatureBlob []byte `gorm:"type:longblob; not null;"`
}

type FeatureDBNew struct {
	ID          int    `gorm:"primary_key"`
	Emp_no      string `gorm:"not null;"`
	FeatureBlob []byte `gorm:"type:longblob; not null;"`
}
