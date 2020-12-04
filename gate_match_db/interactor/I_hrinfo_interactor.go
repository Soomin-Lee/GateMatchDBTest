package interactor

type IHRInfoInteractor interface {
	Create(interface{}) error
	Delete(interface{}) error
	Find(interface{}) ([]interface{}, error)
	GetList(interface{}) ([]interface{}, error)
}
