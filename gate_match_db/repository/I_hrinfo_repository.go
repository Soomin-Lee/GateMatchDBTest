package repository

import "github.com/AlcheraInc/gate_match_db/serializer"

type IHRInfoRepository interface {
	Create(data interface{}) error
	Delete(data interface{}) error
	Find(data interface{}) ([]serializer.SejongEmpInfo, error)
	GetList(data interface{}) ([]serializer.SejongEmpInfo, error)
}
