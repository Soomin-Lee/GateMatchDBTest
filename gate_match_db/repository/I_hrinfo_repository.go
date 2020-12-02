package repository

type HRInfoRepository interface {
	Create(hrinfo interface{}) (interface{}, error)
	Find(id uint) (interface{}, error)
}