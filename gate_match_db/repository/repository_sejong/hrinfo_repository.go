package repository_sejong

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/AlcheraInc/gate_match_db/repository"
	"github.com/AlcheraInc/gate_match_db/serializer"
)

type HRInfoRepository struct {
	db *gorm.DB
}

func NewHRInfoRepository(db *gorm.DB) repository.IHRInfoRepository {
	return &HRInfoRepository{db}
}

func (hr *HRInfoRepository) Create(data interface{}) error {
	newInfo, ok := data.(*entity.EmpInfo)
	if !ok {
		return errors.New("Type error")
	}

	// Create new row in DB
	err := hr.db.Create(&newInfo).Error
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return err
}

func (hr *HRInfoRepository) Delete(data interface{}) error {
	delInfo, ok := data.(*entity.EmpInfo)
	if !ok {
		return errors.New("Type error")
	}

	// Delete row in DB
	err := hr.db.Model(&delInfo).Where("emp_no=?", delInfo.Emp_no).Delete(&delInfo).Error
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (hr *HRInfoRepository) Find(data interface{}) ([]serializer.SejongEmpInfo, error) {
	findInfo, ok := data.(*entity.EmpInfo)
	if !ok {
		return nil, errors.New("Type error")
	}

	// Delete row in DB
	var foundInfos []entity.EmpInfo

	// Find row in DB
	err := hr.db.Model(findInfo).Where("emp_no=?", findInfo.Emp_no).Find(&foundInfos).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var retInfos []serializer.SejongEmpInfo
	for idx := range foundInfos {
		retInfo := serializer.SejongEmpInfo{
			Emp_no:       foundInfos[idx].Emp_no,
			Emp_nm:       foundInfos[idx].Emp_nm,
			Card_no:      foundInfos[idx].Card_no,
			Comp_cd:      foundInfos[idx].Comp_cd,
			Comp_nm:      foundInfos[idx].Comp_nm,
			Dept_cd:      foundInfos[idx].Dept_cd,
			Dept_nm:      foundInfos[idx].Dept_nm,
			Posi_nm:      foundInfos[idx].Posi_nm,
			Regstatus_cd: foundInfos[idx].Regstatus_cd,
			Emptype_cd:   foundInfos[idx].Emptype_cd,
		}
		retInfos = append(retInfos, retInfo)
	}

	return retInfos, nil
}

func (hr *HRInfoRepository) GetList(data interface{}) ([]serializer.SejongEmpInfo, error) {
	listInfo, ok := data.(*entity.FeatureDBNew)
	if !ok {
		return nil, errors.New("Type error")
	}

	// Delete row in DB
	var foundInfos []entity.EmpInfo

	err := hr.db.Model(&listInfo).Find(&foundInfos).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var retInfos []serializer.SejongEmpInfo
	for idx := range foundInfos {
		retInfo := serializer.SejongEmpInfo{
			Emp_no:       foundInfos[idx].Emp_no,
			Emp_nm:       foundInfos[idx].Emp_nm,
			Card_no:      foundInfos[idx].Card_no,
			Comp_cd:      foundInfos[idx].Comp_cd,
			Comp_nm:      foundInfos[idx].Comp_nm,
			Dept_cd:      foundInfos[idx].Dept_cd,
			Dept_nm:      foundInfos[idx].Dept_nm,
			Posi_nm:      foundInfos[idx].Posi_nm,
			Regstatus_cd: foundInfos[idx].Regstatus_cd,
			Emptype_cd:   foundInfos[idx].Emptype_cd,
		}
		retInfos = append(retInfos, retInfo)
	}

	return retInfos, nil
}
