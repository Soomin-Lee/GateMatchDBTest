package repository

import "github.com/AlcheraInc/gate_match_db/serializer"

type IFeatureRepository interface {
	Create(data interface{}) error
	Delete(data interface{}) error
	GetList(data interface{}) ([]serializer.SejongFeatureDBNew, error)
}
