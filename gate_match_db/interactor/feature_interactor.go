package interactor

import (
	"errors"

	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/AlcheraInc/gate_match_db/repository"
	"github.com/AlcheraInc/gate_match_db/serializer"
	utils "github.com/AlcheraInc/go/utils"
)

type FeatureInteractor struct {
	FeatureRepository repository.IFeatureRepository
}

func NewFeatureInteractor(r repository.IFeatureRepository) IFeatureInteractor {
	return &FeatureInteractor{r}
}

func (fi *FeatureInteractor) Create(reqData interface{}) error {
	switch reqData.(type) {
	case serializer.FeatureDB:
		fdata, ok := reqData.(serializer.FeatureDB)
		if !ok {
			return errors.New("Type Error")
		} else {
			entityCreate := &entity.FeatureDB{}
			entityCreate.Emp_no = fdata.Emp_no
			var featureVector [512]float32
			copy(featureVector[:], fdata.FeatureVector[:512])
			feature_bytes, _ := utils.Float32SliceAsByteSlice(featureVector)
			entityCreate.FeatureBlob = feature_bytes
			err := fi.FeatureRepository.Create(entityCreate)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("Type Error")
	}

	return nil
}

func (fi *FeatureInteractor) Delete(reqData interface{}) error {
	switch reqData.(type) {
	case serializer.FeatureDB:
		fdata, ok := reqData.(serializer.FeatureDB)
		if !ok {
			return errors.New("Type Error")
		} else {
			entityDelete := &entity.FeatureDB{}
			entityDelete.Emp_no = fdata.Emp_no
			err := fi.FeatureRepository.Delete(entityDelete)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("Type Error")
	}

	return nil
}

func (fi *FeatureInteractor) GetList(reqData interface{}) ([]interface{}, error) {
	var retData []interface{}

	switch reqData.(type) {
	case serializer.FeatureDB:
		entityFind := &entity.FeatureDB{}
		featureList, err := fi.FeatureRepository.GetList(entityFind)
		if err != nil {
			return nil, err
		}
		retData = make([]interface{}, len(featureList))
		for idx := range featureList {
			retData[idx] = featureList[idx]
		}
	default:
		return nil, errors.New("Type Error")
	}

	return retData, nil
}
