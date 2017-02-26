package service

import (
	m "medigo.com/model"
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
			log.Fatal(err)
		}
	}
}

func StoreDrug(drug m.Drug) {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	err := collection.Insert(drug)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(drug.Cis, "stored")
}

func StoreDrugsOrigins(drugOrigins map[string]float32) {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugsorigins")
	err := collection.Insert(drugOrigins)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("drugOrigins stored")
}
func StoreDrugsDates(drugdates map[string]float32) {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugsdates")
	err := collection.Insert(drugdates)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("drugdates stored")
}

func GetDrugsOriginsFromBase() []m.Titulaire {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	var result []m.Titulaire
	err := collection.Find(bson.M{}).Select(bson.M{"titulaire": 1}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetDrugMarketingDatesFromBase() []m.DateMiseSurLeMarche {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	var result []m.DateMiseSurLeMarche
	err := collection.Find(bson.M{}).Select(bson.M{"datemisesurlemarche": 1}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetDrugsOrigins() []interface{} {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugsorigins")
	result :=  make([]interface{}, 0)
	err := collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetDrugMarketingDates() []interface{} {
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugsdates")
	result :=  make([]interface{}, 0)
	err := collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
