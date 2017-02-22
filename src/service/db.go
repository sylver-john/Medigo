package service

import (
	m "../model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func connectToDb() *mgo.Session {
	session, err := mgo.Dial("****")
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

func GetDrugsorigins() []m.Titulaire {
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
