package registry

import (
	"github.com/AlcheraInc/gate_match_db/interactor"
	"github.com/AlcheraInc/gate_match_db/repository"
	"github.com/AlcheraInc/gate_match_db/repository/repository_sejong"
	"github.com/jinzhu/gorm"
)

type RegistrySejong struct {
	db *gorm.DB
}

func NewRegistrySejong(db *gorm.DB) Registry {
	return &RegistrySejong{db}
}

func (r *RegistrySejong) NewFeatureRepository() repository.IFeatureRepository {
	return repository_sejong.NewFeatureRepository(r.db)
}

func (r *RegistrySejong) NewFeatureInteractor(rfr repository.IFeatureRepository) interactor.IFeatureInteractor {
	return interactor.NewFeatureInteractor(rfr)
}
