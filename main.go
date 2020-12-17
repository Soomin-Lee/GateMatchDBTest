package main

import (
	"log"
	"os"
	"strconv"

	"github.com/AlcheraInc/gate_match_db/database_manager"
	"github.com/AlcheraInc/gate_match_db/feature_db"
	"github.com/AlcheraInc/gate_match_db/migrations"
	"github.com/AlcheraInc/gate_match_db/registry"
	"github.com/jinzhu/gorm"
)

var dbManager database_manager.DatabaseManager

func main() {
	db, err := newDatabaseConnection()
	if err != nil {
		log.Fatalln(err)
		return
	}

	registry := registry.NewRegistry(db)
	featureRepository := registry.NewFeatureRepository()
	featureDB := registry.NewFeatureDB(featureRepository)

	// newFeatureRow := serializer.SejongFeatureDBNew{}
	// newFeatureRow.Emp_no = "Soomin5"
	// newFeatureRow.FeatureVector = make([]float32, 512)
	// for i := range newFeatureRow.FeatureVector {
	// 	newFeatureRow.FeatureVector[i] = 0.5
	// }

	// _, err = featureInteractor.CreateFeatureDBRow(newFeatureRow)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	// newFeatureRow := serializer.SejongFeatureDBNew{}
	// newFeatureRow.Emp_no = "Soomin5"

	// err = featureInteractor.Delete(newFeatureRow)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	err = featureDB.LoadFeatureDB()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Phase 1")
	showMemoryDB(featureDB.MemoryDB)

	fv := make([]float32, 512)
	for i := range fv {
		fv[i] = 0.5
	}
	err = featureDB.CreateFeatureRow("Soomin1", fv)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Phase 2")
	showMemoryDB(featureDB.MemoryDB)

	// for idx := range fr {
	// 	feature, _ := fr[idx].(serializer.SejongFeatureDBNew)
	// 	featureList = append(featureList, feature)
	// 	// nf := serializer.SejongFeatureDBNew{
	// 	// Emp_no:        (string)(elements.Field(0).Interface()),
	// 	// FeatureVector: elements.Field(1).Interface(),
	// 	// }
	// }

	return
}

func showMemoryDB(db []feature_db.FeatureRow) {
	for idx := range db {
		log.Println(db[idx].UID)
	}
}

func newDatabaseConnection() (db *gorm.DB, err error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	portNum, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		return
	}

	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PW")
	db, err = dbManager.ConnectDB(host, int(portNum), dbName, user, password)
	if err != nil {
		return
	}

	err = dbManager.Migrate(migrations.FeatureDBMigrations)
	return
}
