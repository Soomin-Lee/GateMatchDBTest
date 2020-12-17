package registry

import (
	"github.com/AlcheraInc/gate_match_db/feature_db"
	"github.com/AlcheraInc/gate_match_db/repository"
)

type IRegistry interface {
	NewFeatureDB(rfr repository.IFeatureRepository) feature_db.FeatureDB
	NewFeatureRepository() repository.IFeatureRepository
}
