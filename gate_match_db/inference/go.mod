module github.com/AlcheraInc/gate_match_db/inference

go 1.15

require (
	github.com/AlcheraInc/InferenceMsg/InferenceService v0.0.0
)

replace (
	github.com/AlcheraInc/InferenceMsg/InferenceService => ./InferenceMsg/go/github.com/AlcheraInc/InferenceMsg/InferenceService
)