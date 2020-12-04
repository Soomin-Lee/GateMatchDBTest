package entity

type FeatureDB struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"not null;"`
	FeatureBlob []byte `gorm:"type:longblob; not null;"`
}
