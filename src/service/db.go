package service

import (
	m "../model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func connectToDb() *mgo.Session {
	conf := GetConf()
	session, err := mgo.Dial("mongodb://" + conf.Username + ":" + conf.Password + "@" + conf.Host + "/" + conf.Database)
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func StoreDrugCollection(drugs []m.Drug) {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	for i := 0; i < len(drugs); i++ {
		err := collection.Insert(drugs[i])
		if err != nil {
			log.Println(err)
		}
	}
}

func StoreDrug(drug m.Drug) {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	err := collection.Insert(drug)
	if err != nil {
		log.Println(err)
	}
}

func GetDrugsOrigins() []m.Titulaire {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	var result []m.Titulaire
	err := collection.Find(bson.M{}).Select(bson.M{"titulaire": 1}).All(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

func GetDrugMarketingDates() []m.DateMiseSurLeMarche {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	var result []m.DateMiseSurLeMarche
	err := collection.Find(bson.M{}).Select(bson.M{"datemisesurleMarche": 1}).All(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

// func CheckDrugAlreadyInserted(collection *mgo.Collection, cis string) {
// 	err := collection.Find(bson.M{}).
// }
