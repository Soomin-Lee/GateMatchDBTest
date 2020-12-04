package registry

import (
	"github.com/AlcheraInc/gate_match_db/interactor"
	"github.com/AlcheraInc/gate_match_db/repository"
)

type IRegistry interface {
	NewFeatureInteractor(rfr repository.IFeatureRepository) interactor.IFeatureInteractor
	NewFeatureRepository() repository.IFeatureRepository
}
