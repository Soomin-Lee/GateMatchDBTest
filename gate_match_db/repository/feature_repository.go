package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/AlcheraInc/gate_match_db/entity"
)

type FeatureRepository struct {
	db *gorm.DB
}

func NewFeatureRepository(db *gorm.DB) IFeatureRepository {
	return &FeatureRepository{db}
}

func (fr *FeatureRepository) Create(data interface{}) error {
	newFeature, ok := data.(*entity.FeatureRow)
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
	delFeature, ok := data.(*entity.FeatureRow)
	if !ok {
		return errors.New("Type error")
	}

	// Delete row in DB
	err := fr.db.Model(&delFeature).Where("uid=?", delFeature.UID).Delete(&delFeature).Error
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (fr *FeatureRepository) Find(data interface{}) ([]entity.FeatureRow, error) {
	findFeature, ok := data.(*entity.FeatureRow)
	if !ok {
		return nil, errors.New("Type error")
	}

	// Find row in DB
	var foundFeatures []entity.FeatureRow

	err := fr.db.Model(&findFeature).Where("uid=?", findFeature.UID).Find(&foundFeatures).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// var retFeatures []serializer.FeatureDB
	// for idx := range foundFeatures {
	// 	featureName := foundFeatures[idx].Name
	// 	featureBytes := foundFeatures[idx].FeatureBlob
	// 	featureVector := utils.ByteSliceAsFloat32Slice(featureBytes, len(featureBytes))
	// 	retFeature := serializer.FeatureDB{
	// 		Name:          featureName,
	// 		FeatureVector: featureVector,
	// 	}
	// 	retFeatures = append(retFeatures, retFeature)
	// }

	return foundFeatures, nil
}

func (fr *FeatureRepository) GetList(data interface{}) ([]entity.FeatureRow, error) {
	listFeature, ok := data.(*entity.FeatureRow)
	if !ok {
		return nil, errors.New("Type error")
	}

	var foundFeatures []entity.FeatureRow

	err := fr.db.Model(&listFeature).Find(&foundFeatures).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return foundFeatures, nil

	// var retFeatures []serializer.FeatureDB
	// for idx := range foundFeatures {
	// 	featureName := foundFeatures[idx].Name
	// 	featureBytes := foundFeatures[idx].FeatureBlob
	// 	featureVector := utils.ByteSliceAsFloat32Slice(featureBytes, len(featureBytes))
	// 	retFeature := serializer.FeatureDB{
	// 		Name:          featureName,
	// 		FeatureVector: featureVector,
	// 	}
	// 	retFeatures = append(retFeatures, retFeature)
	// }

	// return retFeatures, nil
}
