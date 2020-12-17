package feature_db

type FeatureRow struct {
	UID           string
	FeatureVector [512]float32
}
