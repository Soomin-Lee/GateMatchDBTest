package repository_sejong

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/AlcheraInc/gate_match_db/repository"
	"github.com/AlcheraInc/gate_match_db/serializer"
	"github.com/AlcheraInc/go/utils"
)

type FeatureRepository struct {
	db *gorm.DB
}

func NewFeatureRepository(db *gorm.DB) repository.IFeatureRepository {
	return &FeatureRepository{db}
}

func (fr *FeatureRepository) Create(data interface{}) error {
	newFeature, ok := data.(*entity.FeatureDBNew)
	if !ok {
		return errors.New("Type error")
	}

	// Create new row in DB
	err := fr.db.Create(&newFeature).Error
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return err
}

func (fr *FeatureRepository) Delete(data interface{}) error {
	delFeature, ok := data.(*entity.FeatureDBNew)
	if !ok {
		return errors.New("Type error")
	}

	// Delete row in DB
	err := fr.db.Where("emp_no=?", delFeature.Emp_no).Delete(&delFeature).Error
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (fr *FeatureRepository) GetList(data interface{}) ([]serializer.SejongFeatureDBNew, error) {
	_, ok := data.(*entity.FeatureDBNew)
	if !ok {
		return nil, errors.New("Type error")
	}

	// Delete row in DB
	var foundFeatures []entity.FeatureDBNew

	err := fr.db.Find(&foundFeatures).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var retFeatures []serializer.SejongFeatureDBNew
	for idx := range foundFeatures {
		featureEmpNo := foundFeatures[idx].Emp_no
		featureBytes := foundFeatures[idx].FeatureBlob
		featureVector := utils.ByteSliceAsFloat32Slice(featureBytes, len(featureBytes))
		retFeature := serializer.SejongFeatureDBNew{
			Emp_no:        featureEmpNo,
			FeatureVector: featureVector,
		}
		retFeatures = append(retFeatures, retFeature)
	}

	return retFeatures, nil
}
