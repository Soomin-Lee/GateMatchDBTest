module github.com/AlcheraInc/gate_match_db/migrations

go 1.15

require (
    github.com/AlcheraInc/gate_match_db/entity v0.0.0
)

replace (
    github.com/AlcheraInc/gate_match_db/entity => ../entity
)