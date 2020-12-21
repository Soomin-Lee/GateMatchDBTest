module github.com/AlcheraInc/GateMatchDBTest

go 1.15

require (
	github.com/AlcheraInc/InferenceMsg/InferenceService v0.0.0
	github.com/AlcheraInc/gate_match_db/database_manager v0.0.0
	github.com/AlcheraInc/gate_match_db/entity v0.0.0
	github.com/AlcheraInc/gate_match_db/feature_db v0.0.0
	github.com/AlcheraInc/gate_match_db/inference v0.0.0
	github.com/AlcheraInc/gate_match_db/registry v0.0.0
	github.com/AlcheraInc/gate_match_db/repository v0.0.0
	github.com/AlcheraInc/go v0.0.0-20201130060314-41585e1a050a // indirect
	github.com/jinzhu/gorm v1.9.16
	google.golang.org/grpc v1.34.0 // indirect
	gopkg.in/gormigrate.v1 v1.6.0 // indirect
)

replace (
	github.com/AlcheraInc/InferenceMsg/InferenceService => ./gate_match_db/inference/InferenceMsg/go/github.com/AlcheraInc/InferenceMsg/InferenceService
	github.com/AlcheraInc/gate_match_db/database_manager => ./gate_match_db/database_manager
	github.com/AlcheraInc/gate_match_db/entity => ./gate_match_db/entity
	github.com/AlcheraInc/gate_match_db/feature_db => ./gate_match_db/feature_db
	github.com/AlcheraInc/gate_match_db/inference => ./gate_match_db/inference
	github.com/AlcheraInc/gate_match_db/registry => ./gate_match_db/registry
	github.com/AlcheraInc/gate_match_db/repository => ./gate_match_db/repository
)
