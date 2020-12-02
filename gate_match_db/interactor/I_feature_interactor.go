package interactor

type IFeatureInteractor interface {
	CreateFeatureDBRow(interface{}) (interface{}, error)
}
