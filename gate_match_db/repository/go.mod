module github.com/AlcheraInc/gate_match_db/repository

go 1.15

require (
    github.com/AlcheraInc/gate_match_db/entity v0.0.0
    github.com/AlcheraInc/gate_match_db/serializer v0.0.0
)

replace (
    github.com/AlcheraInc/gate_match_db/entity => ../entity
    github.com/AlcheraInc/gate_match_db/serializer => ../serializer
)