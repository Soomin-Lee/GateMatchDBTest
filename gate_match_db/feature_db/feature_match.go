package feature_db

import (
	"errors"
	"log"
	"math"
	"sync"
)

type DistanceResult struct {
	UID      string
	Distance float32
}

// 단일 Result 반환
func (fi *FeatureDB) MatchFeature(fv []float32, distanceThreshold float32) (DistanceResult, error) {
	if fi.MemoryDB == nil {
		return DistanceResult{}, errors.New("Feature DB Not ready")
	}

	var requestFeatureVector [512]float32
	copy(requestFeatureVector[:], fv[:512])

	// Match start
	wg := sync.WaitGroup{}
	wg.Add(2)

	var minA float32
	var minB float32
	var distResultA DistanceResult
	var distResultB DistanceResult

	fun := func(arr []FeatureRow, min *float32, distResult *DistanceResult) {
		*min = math.MaxFloat32
		for _, row := range arr {
			distance := FeatureDistanceFloat32(row.FeatureVector, requestFeatureVector)
			if distance < distanceThreshold && distance < *min {
				*min = distance
				*&distResult.UID = row.UID
				*&distResult.Distance = distance
			}
		}
		wg.Done()
	}

	dbLength := len(fi.MemoryDB)
	log.Println("DB Length :", dbLength)
	go fun(fi.MemoryDB[:dbLength/2], &minA, &distResultA)
	go fun(fi.MemoryDB[dbLength/2:], &minB, &distResultB)
	wg.Wait()

	var result DistanceResult
	var resultMinDistance float32
	if minA < minB {
		result = distResultA
		resultMinDistance = minA
	} else {
		result = distResultB
		resultMinDistance = minB
	}

	if resultMinDistance == math.MaxFloat32 {
		result.UID = "-1"
		result.Distance = -1.0
		log.Println("Cannot find face")
	}

	return result, nil
}

// Threshold 이내의 모든 Result 반환 (passed_empno와 같음)
func (fi *FeatureDB) MatchFeatureAll(fv []float32, distanceThreshold float32) (DistanceResult, []DistanceResult, error) {
	if fi.MemoryDB == nil {
		return DistanceResult{}, nil, errors.New("Feature DB Not ready")
	}

	var requestFeatureVector [512]float32
	copy(requestFeatureVector[:], fv[:512])

	// Match start
	wg := sync.WaitGroup{}
	wg.Add(2)

	var passed_uids_a []DistanceResult
	var passed_uids_b []DistanceResult

	var minA float32
	var minB float32
	var distResultA DistanceResult
	var distResultB DistanceResult

	fun := func(arr []FeatureRow, min *float32, distResult *DistanceResult, passed *[]DistanceResult) {
		*min = math.MaxFloat32
		for _, row := range arr {
			distance := FeatureDistanceFloat32(row.FeatureVector, requestFeatureVector)
			if distance < distanceThreshold {
				dr := DistanceResult{}
				dr.UID = row.UID
				dr.Distance = distance
				*passed = append(*passed, dr)
			}
			if distance < distanceThreshold && distance < *min {
				*min = distance
				*&distResult.UID = row.UID
				*&distResult.Distance = distance
			}
		}
		wg.Done()
	}

	dbLength := len(fi.MemoryDB)
	log.Println("DB Length :", dbLength)
	go fun(fi.MemoryDB[:dbLength/2], &minA, &distResultA, &passed_uids_a)
	go fun(fi.MemoryDB[dbLength/2:], &minB, &distResultB, &passed_uids_b)
	wg.Wait()

	var passed_uids []DistanceResult
	passed_uids = append(passed_uids, passed_uids_a...)
	passed_uids = append(passed_uids, passed_uids_b...)

	var result DistanceResult
	var resultMinDistance float32
	if minA < minB {
		result = distResultA
		resultMinDistance = minA
	} else {
		result = distResultB
		resultMinDistance = minB
	}

	if resultMinDistance == math.MaxFloat32 {
		result.UID = "-1"
		result.Distance = -1.0
		log.Println("Cannot find face")
	}

	return result, passed_uids, nil
}

func FeatureDistanceFloat32(fa [512]float32, fb [512]float32) float32 {
	var dp float32 = 0.0
	for idx := range fa {
		dp += fa[idx] * fb[idx]
	}
	return (1 - dp) / 2.0
}
