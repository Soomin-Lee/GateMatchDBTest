module github.com/AlcheraInc/gate_match_db/interactor

go 1.15

require (
    github.com/AlcheraInc/gate_match_db/serializer v0.0.0
    github.com/AlcheraInc/gate_match_db/repository v0.0.0
)

replace (
    github.com/AlcheraInc/gate_match_db/serializer => ../serializer
    github.com/AlcheraInc/gate_match_db/repository => ../repository
)