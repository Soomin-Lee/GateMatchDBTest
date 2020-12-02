package repository

type IFeatureRepository interface {
	Create(feature interface{}) (interface{}, error)
}
