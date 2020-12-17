package registry

import (
	"github.com/AlcheraInc/gate_match_db/feature_db"
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

func (r *Registry) NewFeatureDB(rfr repository.IFeatureRepository) feature_db.FeatureDB {
	return feature_db.NewFeatureDB(rfr)
}
