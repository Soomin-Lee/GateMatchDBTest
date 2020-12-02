package repository_sejong

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/AlcheraInc/gate_match_db/repository"
)

type FeatureRepository struct {
	db *gorm.DB
}

func NewFeatureRepository(db *gorm.DB) repository.IFeatureRepository {
	return &FeatureRepository{db}
}

func (fr *FeatureRepository) Create(feature interface{}) (interface{}, error) {
	newFeature, ok := feature.(*entity.FeatureDBNew)
	if !ok {
		log.Fatalln("NO")
		return nil, errors.New("Type error")
	}

	// Create new row in DB
	err := fr.db.Create(&newFeature).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return newFeature, err
}
