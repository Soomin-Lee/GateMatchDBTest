package main

import (
	"os"
	"strconv"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/AlcheraInc/gate_match_db/registry"
	"github.com/AlcheraInc/gate_match_db/database_manager"
	"github.com/AlcheraInc/gate_match_db/migrations"
	"github.com/AlcheraInc/gate_match_db/serializer"
)

var dbManager database_manager.DatabaseManager

func main() {
	db, err := newDatabaseConnection()
	if err != nil {
		log.Fatalln(err)
		return
	}

	registry := registry.NewRegistrySejong(db)
	featureRepository := registry.NewFeatureRepository()
	featureInteractor := registry.NewFeatureInteractor(featureRepository)
	
	newFeatureRow := serializer.SejongFeatureDBNew{}
	newFeatureRow.Emp_no = "Soomin"
	newFeatureRow.FeatureVector = make([]float32, 512)
	for i := range newFeatureRow.FeatureVector {
		newFeatureRow.FeatureVector[i] = 0.5
	}

	_, err = featureInteractor.CreateFeatureDBRow(newFeatureRow)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
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

	err = dbManager.Migrate(migrations.SejongFeatureMigrations)
	return
}
