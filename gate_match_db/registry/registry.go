package registry

import (
	"github.com/AlcheraInc/gate_match_db/interactor"
	"github.com/AlcheraInc/gate_match_db/repository"
	"github.com/jinzhu/gorm"
)

type Registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) IRegistry {
	return &Registry{db}
}

func (r *Registry) NewFeatureRepository() repository.IFeatureRepository {
	return repository.NewFeatureRepository(r.db)
}

func (r *Registry) NewFeatureInteractor(rfr repository.IFeatureRepository) interactor.IFeatureInteractor {
	return interactor.NewFeatureInteractor(rfr)
}
