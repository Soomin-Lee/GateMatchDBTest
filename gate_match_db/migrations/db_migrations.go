package migrations

import (
	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var FeatureDBMigrations = []*gormigrate.Migration{
	{
		ID: "202012041100",
		Migrate: func(tx *gorm.DB) (err error) {
			featuredb := new(entity.FeatureDB)
			tx.LogMode(true)

			err = tx.AutoMigrate(featuredb).Error
			tx.LogMode(false)

			return
		},
		Rollback: func(tx *gorm.DB) (err error) {
			featuredb := new(entity.FeatureDB)
			tx.LogMode(true)

			err = tx.DropTableIfExists(featuredb).Error
			return
		},
	},
}
