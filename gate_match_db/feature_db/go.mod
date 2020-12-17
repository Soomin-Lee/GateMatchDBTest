module github.com/AlcheraInc/gate_match_db/feature_db

go 1.15

require (
    github.com/AlcheraInc/gate_match_db/repository v0.0.0
)

replace (
    github.com/AlcheraInc/gate_match_db/repository => ../repository
)