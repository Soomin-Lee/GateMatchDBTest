package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/AlcheraInc/gate_match_db/serializer"
	"github.com/AlcheraInc/go/utils"
)

type FeatureRepository struct {
	db *gorm.DB
}

func NewFeatureRepository(db *gorm.DB) IFeatureRepository {
	return &FeatureRepository{db}
}

func (fr *FeatureRepository) Create(data interface{}) error {
	newFeature, ok := data.(*entity.FeatureDB)
	if !ok {
		return errors.New("Type error")
	}

	// Create new row in DB
	err := fr.db.Model(&newFeature).Create(&newFeature).Error
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return err
}

func (fr *FeatureRepository) Delete(data interface{}) error {
	delFeature, ok := data.(*entity.FeatureDB)
	if !ok {
		return errors.New("Type error")
	}

	// Delete row in DB
	err := fr.db.Model(&delFeature).Where("name=?", delFeature.Name).Delete(&delFeature).Error
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (fr *FeatureRepository) GetList(data interface{}) ([]interface{}, error) {
	listFeature, ok := data.(*entity.FeatureDB)
	if !ok {
		return nil, errors.New("Type error")
	}

	// Delete row in DB
	var foundFeatures []entity.FeatureDB

	err := fr.db.Model(&listFeature).Find(&foundFeatures).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var retFeatures []serializer.FeatureDB
	for idx := range foundFeatures {
		featureName := foundFeatures[idx].Name
		featureBytes := foundFeatures[idx].FeatureBlob
		featureVector := utils.ByteSliceAsFloat32Slice(featureBytes, len(featureBytes))
		retFeature := serializer.FeatureDB{
			Name:          featureName,
			FeatureVector: featureVector,
		}
		retFeatures = append(retFeatures, retFeature)
	}

	return retFeatures, nil
}
