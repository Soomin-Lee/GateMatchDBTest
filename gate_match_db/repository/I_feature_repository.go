package repository

type IFeatureRepository interface {
	Create(data interface{}) error
	Delete(data interface{}) error
	GetList(data interface{}) ([]interface{}, error)
}
