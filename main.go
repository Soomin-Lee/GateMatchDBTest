package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/AlcheraInc/GateMatchBase/database_manager"
	"github.com/AlcheraInc/GateMatchBase/feature_db"
	"github.com/AlcheraInc/GateMatchBase/inference"
	"github.com/AlcheraInc/GateMatchBase/registry"
	"github.com/jinzhu/gorm"
)

var dbManager database_manager.DatabaseManager

func main() {
	err := inference.ConnectInferenceService()
	if err != nil {
		log.Fatalln(err)
		return
	}

	imgdata, err := ioutil.ReadFile("/home/leesoomin/dev/SolutionTeam/IntegratedFaceServer/testdata/images/feature_test.png")
	// imgdata, err := ioutil.ReadFile("/home/magmatart/Alchera/IntegratedFaceServer/testdata/images/smlee_1.jpg")
	if err != nil {
		log.Fatalln(err)
		return
	}

	featureVector, err := inference.InferenceFeatureAlignedImage(imgdata)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for i := range featureVector {
		fmt.Print(featureVector[i], ", ")
	}
	log.Println()

	db, err := newDatabaseConnection()
	if err != nil {
		log.Fatalln(err)
		return
	}

	registry := registry.NewRegistry(db)
	featureRepository := registry.NewFeatureRepository()
	featureDB := registry.NewFeatureDB(featureRepository)

	err = featureDB.LoadFeatureDB()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Phase 1")
	showMemoryDB(featureDB.MemoryDB)

	err = featureDB.CreateFeatureRow("RealFeature", featureVector)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Input Complete")
	showMemoryDB(featureDB.MemoryDB)

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

	// log.Println("Phase 2")
	// showMemoryDB(featureDB.MemoryDB)

	// err = featureDB.DeleteFeatureRow("Soomin0")
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

	// matchResult, passedResults, err := featureDB.MatchFeatureAll(fv, 0.2)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println("UID :", matchResult.UID)
	// log.Println("Distance :", matchResult.Distance)
	// log.Println("Passed Count :", len(passedResults))

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

	err = dbManager.Migrate(feature_db.Migrations)
	return
}
