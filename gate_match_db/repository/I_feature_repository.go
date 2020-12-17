package repository

import "github.com/AlcheraInc/gate_match_db/entity"

type IFeatureRepository interface {
	Create(data interface{}) error
	Delete(data interface{}) error
	Find(data interface{}) ([]entity.FeatureRow, error)
	GetList(data interface{}) ([]entity.FeatureRow, error)
}
