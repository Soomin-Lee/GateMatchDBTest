package inference

import (
	"errors"

	InferenceMsg "github.com/AlcheraInc/InferenceMsg/InferenceService"
)

func InferenceFeatureAlignedImage(imgdata []byte) ([]float32, error){
	var retFeatureVector []float32
	
	feature_req := InferenceMsg.FeatureRequest{}
	feature_res := &InferenceMsg.FeatureResult{}

	if len(imgdata) < 1 {
		return retFeatureVector, errors.New("Image data not valid")
	}

	feature_req.Imgdata = imgdata

	feature_res, err := inferenceClient.RunFeature(grpcContext, &feature_req)
	if err != nil {
		return retFeatureVector, err
	}
	if feature_res.ReturnMsg.GetCode() != "INF-00" {
		return retFeatureVector, errors.New(feature_res.ReturnMsg.GetCode())
	}

	retFeatureVector = feature_res.GetFeatureVector()
	return retFeatureVector, nil
}

func InferenceFeatureFullImage(imgdata []byte) ([]float32, error) {
	var retFeatureVector []float32
	
	feature_req := InferenceMsg.FaceFeaturesRequest{}
	feature_res := &InferenceMsg.FaceFeaturesResult{}

	if len(imgdata) < 1 {
		return retFeatureVector, errors.New("Image data not valid")
	}

	feature_req.Imgdata = imgdata

	feature_res, err := inferenceClient.RunFaceFeatures(grpcContext, &feature_req)
	if err != nil {
		return retFeatureVector, err
	}
	if feature_res.ReturnMsg.GetCode() != "INF-00" {
		return retFeatureVector, errors.New(feature_res.ReturnMsg.GetCode())
	}

	if feature_res.GetCount() < 1 {
		return retFeatureVector, errors.New("Cannot find any faces")
	}

	retFeatureVector = feature_res.GetFacesWithFeatures()[0].FeatureVector
	return retFeatureVector, nil
}