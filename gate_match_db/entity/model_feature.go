package entity

type FeatureRow struct {
	ID          int    `gorm:"primary_key"`
	UID         string `gorm:"not null;"`
	FeatureBlob []byte `gorm:"type:longblob; not null;"`
}
