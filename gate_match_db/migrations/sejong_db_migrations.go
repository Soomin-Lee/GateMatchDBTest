package migrations

import (
	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var SejongFeatureMigrations = []*gormigrate.Migration{
	{
		ID: "202012031100",
		Migrate: func(tx *gorm.DB) (err error) {
			featuredb := new(entity.FeatureDBNew)
			tx.LogMode(true)

			err = tx.AutoMigrate(featuredb).Error
			tx.LogMode(false)

			return
		},
		Rollback: func(tx *gorm.DB) (err error) {
			featuredb := new(entity.FeatureDBNew)
			tx.LogMode(true)

			err = tx.DropTableIfExists(featuredb).Error
			return
		},
	},
}
