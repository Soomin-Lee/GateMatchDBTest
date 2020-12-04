module github.com/AlcheraInc/GateMatchDBTest

go 1.15

require (
	github.com/AlcheraInc/gate_match_db/database_manager v0.0.0
	github.com/AlcheraInc/gate_match_db/entity v0.0.0
	github.com/AlcheraInc/gate_match_db/interactor v0.0.0
	github.com/AlcheraInc/gate_match_db/migrations v0.0.0
	github.com/AlcheraInc/gate_match_db/registry v0.0.0
	github.com/AlcheraInc/gate_match_db/repository v0.0.0
	github.com/AlcheraInc/gate_match_db/serializer v0.0.0
	github.com/AlcheraInc/go v0.0.0-20201130060314-41585e1a050a // indirect
	github.com/jinzhu/gorm v1.9.16
	gopkg.in/gormigrate.v1 v1.6.0 // indirect
)

replace (
	github.com/AlcheraInc/gate_match_db/database_manager => ./gate_match_db/database_manager
	github.com/AlcheraInc/gate_match_db/entity => ./gate_match_db/entity
	github.com/AlcheraInc/gate_match_db/interactor => ./gate_match_db/interactor
	github.com/AlcheraInc/gate_match_db/migrations => ./gate_match_db/migrations
	github.com/AlcheraInc/gate_match_db/registry => ./gate_match_db/registry
	github.com/AlcheraInc/gate_match_db/repository => ./gate_match_db/repository
	github.com/AlcheraInc/gate_match_db/serializer => ./gate_match_db/serializer
)
