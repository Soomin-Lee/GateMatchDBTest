package interactor

import (
	"errors"

	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/AlcheraInc/gate_match_db/repository"
	"github.com/AlcheraInc/gate_match_db/serializer"
)

type HRInfoInteractor struct {
	HRInfoRepository repository.IHRInfoRepository
}

func NewHRInfoInteractor(r repository.IHRInfoRepository) IHRInfoInteractor {
	return &HRInfoInteractor{r}
}

func (hi *HRInfoInteractor) Create(reqData interface{}) error {
	switch reqData.(type) {
	case serializer.SejongEmpInfo:
		fdata, ok := reqData.(serializer.SejongEmpInfo)
		if !ok {
			return errors.New("Type Error")
		} else {
			entityCreate := &entity.EmpInfo{
				Emp_no:       fdata.Emp_no,
				Emp_nm:       fdata.Emp_nm,
				Card_no:      fdata.Card_no,
				Comp_cd:      fdata.Comp_cd,
				Comp_nm:      fdata.Comp_nm,
				Dept_cd:      fdata.Dept_cd,
				Dept_nm:      fdata.Dept_nm,
				Posi_nm:      fdata.Posi_nm,
				Regstatus_cd: fdata.Regstatus_cd,
				Emptype_cd:   fdata.Emptype_cd,
			}

			err := hi.HRInfoRepository.Create(entityCreate)
			if err != nil {
				return err
			}
		}
	case serializer.ParadiseEmpInfo:
		// Paradise DB에 맞춰 적용
		break
	default:
		return errors.New("Type Error")
	}

	return nil
}

func (hi *HRInfoInteractor) Delete(reqData interface{}) error {
	switch reqData.(type) {
	case serializer.SejongEmpInfo:
		fdata, ok := reqData.(serializer.SejongEmpInfo)
		if !ok {
			return errors.New("Type Error")
		} else {
			entityDelete := &entity.EmpInfo{}
			entityDelete.Emp_no = fdata.Emp_no
			err := hi.HRInfoRepository.Delete(entityDelete)
			if err != nil {
				return err
			}
		}
	case serializer.ParadiseEmpInfo:
		break
	default:
		return errors.New("Type Error")
	}

	return nil
}

func (hi *HRInfoInteractor) Find(reqData interface{}) ([]interface{}, error) {
	var retData []interface{}

	switch reqData.(type) {
	case serializer.SejongEmpInfo:
		fdata, ok := reqData.(serializer.SejongEmpInfo)
		if !ok {
			return nil, errors.New("Type Error")
		} else {
			entityFind := &entity.EmpInfo{}
			entityFind.Emp_no = fdata.Emp_no
			infoList, err := hi.HRInfoRepository.Find(entityFind)
			if err != nil {
				return nil, err
			}
			retData = make([]interface{}, len(infoList))
			for idx := range infoList {
				retData[idx] = infoList[idx]
			}
		}
	case serializer.ParadiseFeatureDB:
		break
	default:
		return nil, errors.New("Type Error")
	}

	return retData, nil
}

func (hi *HRInfoInteractor) GetList(reqData interface{}) ([]interface{}, error) {
	var retData []interface{}

	switch reqData.(type) {
	case serializer.SejongEmpInfo:
		entityFind := &entity.EmpInfo{}
		infoList, err := hi.HRInfoRepository.GetList(entityFind)
		if err != nil {
			return nil, err
		}
		retData = make([]interface{}, len(infoList))
		for idx := range infoList {
			retData[idx] = infoList[idx]
		}
	case serializer.ParadiseFeatureDB:
		break
	default:
		return nil, errors.New("Type Error")
	}

	return retData, nil
}
