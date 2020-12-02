package interactor

import (
	"errors"
	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/AlcheraInc/gate_match_db/serializer"
	"github.com/AlcheraInc/gate_match_db/repository"
	utils "github.com/AlcheraInc/go/utils"
)

type FeatureInteractor struct {
	FeatureRepository repository.IFeatureRepository
}

func NewFeatureInteractor(r repository.IFeatureRepository) IFeatureInteractor {
	return &FeatureInteractor{r}
}

func (fi *FeatureInteractor) CreateFeatureDBRow(reqData interface{}) (interface{}, error) {
	var retData interface{} 
	switch reqData.(type) {
	case serializer.SejongFeatureDBNew:
		fdata, ok := reqData.(serializer.SejongFeatureDBNew)
		if !ok {
			return nil, errors.New("Type Error")
		} else {
			fdbnew := &entity.FeatureDBNew{}
			fdbnew.Emp_no = fdata.Emp_no
			var featureVector [512]float32
			copy(featureVector[:], fdata.FeatureVector[:512])			
			feature_bytes, _ := utils.Float32SliceAsByteSlice(featureVector)
			fdbnew.FeatureBlob = feature_bytes
			fi.FeatureRepository.Create(fdbnew)
			retData = fdbnew
		}
	case serializer.ParadiseFeatureDB:
		// Paradise DB에 맞춰 적용
		break
	default:
		return nil, errors.New("Type Error")
	}
	return retData, nil
}