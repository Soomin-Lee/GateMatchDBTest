module github.com/AlcheraInc/gate_match_db/registry

go 1.15

require (
    github.com/AlcheraInc/gate_match_db/interactor v0.0.0
    github.com/AlcheraInc/gate_match_db/repository v0.0.0
    github.com/AlcheraInc/gate_match_db/repository/repository_sejong v0.0.0
)

replace (
    github.com/AlcheraInc/gate_match_db/interactor => ../interactor
    github.com/AlcheraInc/gate_match_db/repository => ../repository
    github.com/AlcheraInc/gate_match_db/repository/repository_sejong => ../repository/repository_sejong
)