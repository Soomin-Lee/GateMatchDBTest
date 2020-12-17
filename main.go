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

	// for i := 0; i < 10; i++ {
	// 	fv := make([]float32, 512)
	// 	for k := range fv {
	// 		fv[k] = 0.1 * float32(i)
	// 	}
	// 	err = featureDB.CreateFeatureRow("Soomin"+strconv.Itoa(i), fv)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// }

	// err = featureDB.CreateFeatureRow("Soomin1", fv)
	// err = featureDB.CreateFeatureRow("Soomin1", fv)
	// err = featureDB.CreateFeatureRow("Soomin2", fv)
	// err = featureDB.CreateFeatureRow("Soomin2", fv)
	// err = featureDB.CreateFeatureRow("Soomin3", fv)
	// err = featureDB.CreateFeatureRow("Soomin3", fv)
	// err = featureDB.CreateFeatureRow("Soomin3", fv)
	// err = featureDB.CreateFeatureRow("Soomin4", fv)
	// err = featureDB.CreateFeatureRow("Soomin4", fv)
	// err = featureDB.CreateFeatureRow("Soomin4", fv)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	log.Println("Phase 2")
	showMemoryDB(featureDB.MemoryDB)

	fv := make([]float32, 512)
	for k := range fv {
		fv[k] = 0.5
	}

	// err = featureDB.DeleteFeatureRow("Soomin2")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// log.Println("Phase 3")
	// showMemoryDB(featureDB.MemoryDB)

	// matchResult, err := featureDB.MatchFeature(fv, 0.2)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println("UID :", matchResult.UID)
	// log.Println("Distance :", matchResult.Distance)

	matchResult, passedResults, err := featureDB.MatchFeatureAll(fv, 0.2)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("UID :", matchResult.UID)
	log.Println("Distance :", matchResult.Distance)
	log.Println("Passed Count :", len(passedResults))

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
	log.Println(len(db))
	for idx := range db {
		log.Println("UID :", db[idx].UID)
		// for i := 0; i < 512; i++ {
		// 	fmt.Print(db[idx].FeatureVector[i], " ")
		// }
		// fmt.Println()
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
