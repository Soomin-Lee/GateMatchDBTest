package interactor

type IFeatureInteractor interface {
	Create(interface{}) error
	Delete(interface{}) error
	GetList(interface{}) ([]interface{}, error)
}
