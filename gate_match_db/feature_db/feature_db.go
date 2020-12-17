package feature_db

import (
	"errors"

	"github.com/AlcheraInc/gate_match_db/entity"
	"github.com/AlcheraInc/gate_match_db/repository"
	utils "github.com/AlcheraInc/go/utils"
)

type FeatureDB struct {
	FeatureRepository repository.IFeatureRepository
	MemoryDB          []FeatureRow
}

func NewFeatureDB(r repository.IFeatureRepository) FeatureDB {
	return FeatureDB{r, nil}
}

func (fi *FeatureDB) CreateFeatureRow(uid string, fv []float32) error {
	if fi.MemoryDB == nil {
		return errors.New("Feature DB Not ready")
	}

	entityCreate := &entity.FeatureRow{}
	entityCreate.UID = uid
	var featureVector [512]float32
	copy(featureVector[:], fv[:512])
	feature_bytes, _ := utils.Float32SliceAsByteSlice(featureVector)
	entityCreate.FeatureBlob = feature_bytes
	err := fi.FeatureRepository.Create(entityCreate)
	if err != nil {
		return err
	}

	createRow := FeatureRow{
		UID:           uid,
		FeatureVector: featureVector,
	}
	fi.InsertFeatureRowToMemoryDB(createRow)

	return nil
}

func (fi *FeatureDB) DeleteFeatureRow(uid string) error {
	if fi.MemoryDB == nil {
		return errors.New("Feature DB Not ready")
	}

	entityDelete := &entity.FeatureRow{}
	entityDelete.UID = uid
	err := fi.FeatureRepository.Delete(entityDelete)
	if err != nil {
		return err
	}

	delUIDs := fi.FindFeatureIndexByUID(uid)
	fi.DeleteFeatureRowsFromMemoryDB(delUIDs)

	return nil
}

func (fi *FeatureDB) LoadFeatureDB() error {
	if fi.MemoryDB != nil {
		return errors.New("Feature DB Already loaded")
	}

	entityFind := &entity.FeatureRow{}
	featureList, err := fi.FeatureRepository.GetList(entityFind)
	if err != nil {
		return err
	}

	fi.MemoryDB = make([]FeatureRow, 0)

	for idx := range featureList {
		featureUID := featureList[idx].UID
		featureBytes := featureList[idx].FeatureBlob
		featureVector := utils.ByteSliceAsFloat32Slice(featureBytes, len(featureBytes))
		newFeatureRow := FeatureRow{
			UID:           featureUID,
			FeatureVector: featureVector,
		}
		fi.MemoryDB = append(fi.MemoryDB, newFeatureRow)
	}

	return nil
}

func (fi *FeatureDB) InsertFeatureRowToMemoryDB(row FeatureRow) {
	fi.MemoryDB = append(fi.MemoryDB, row)
}

func (fi *FeatureDB) FindFeatureIndexByUID(uid string) []int {
	retUIDs := []int{}
	for idx := range fi.MemoryDB {
		if fi.MemoryDB[idx].UID == uid {
			retUIDs = append(retUIDs, idx)
		}
	}
	return retUIDs
}

func (fi *FeatureDB) DeleteFeatureRowsFromMemoryDB(uids []int) {
	removeElem := func(db []FeatureRow, idx int) []FeatureRow {
		return append(db[:idx], db[idx+1:]...)
	}
	for idx := range uids {
		fi.MemoryDB = removeElem(fi.MemoryDB, uids[idx])
		// sub 1 to remain indices
		for subIdx := idx + 1; subIdx < len(uids); subIdx++ {
			uids[subIdx]--
		}
	}
}
