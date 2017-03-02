package service

import (
	m "Medigo/src/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	log "github.com/cihub/seelog"
)

func connectToDb() *mgo.Session {
	defer log.Flush()
	conf := GetConf()
	session, err := mgo.Dial("mongodb://" + conf.Username + ":" + conf.Password + "@" + conf.Host + "/" + conf.Database)
	if err != nil {
		log.Info(err)
	}
	return session
}

func StoreDrugCollection(drugs []m.Drug) {
	defer log.Flush()
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	for i := 0; i < len(drugs); i++ {
		err := collection.Insert(drugs[i])
		if err != nil {
			log.Info(err)
		}
	}
}

func StoreDrug(drug m.Drug) {
	defer log.Flush()
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	err := collection.Insert(drug)
	if err != nil {
		log.Info(err)
	}
	log.Info(drug.Cis, "stored")
}

func StoreDrugsOrigins(drugOrigins map[string]float32) {
	defer log.Flush()
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugsorigins")
	err := collection.Insert(drugOrigins)
	if err != nil {
		log.Info(err)
	}
	log.Info("drugOrigins stored")
}
func StoreDrugsDates(drugdates map[string]float32) {
	defer log.Flush()
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugsdates")
	err := collection.Insert(drugdates)
	if err != nil {
		log.Info(err)
	}
	log.Info("drugdates stored")
}

func GetDrugsOriginsFromBase() []m.Titulaire {
	defer log.Flush()
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	var result []m.Titulaire
	err := collection.Find(bson.M{}).Select(bson.M{"titulaire": 1}).All(&result)
	if err != nil {
		log.Info(err)
	}
	return result
}

func GetDrugMarketingDatesFromBase() []m.DateMiseSurLeMarche {
	defer log.Flush()
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugs")
	var result []m.DateMiseSurLeMarche
	err := collection.Find(bson.M{}).Select(bson.M{"datemisesurlemarche": 1}).All(&result)
	if err != nil {
		log.Info(err)
	}
	return result
}

func GetDrugsOrigins() []interface{} {

	testConfig := `
	<seelog>
	<outputs>
		<file path="./log/log.log"/>
	</outputs>
	</seelog>
	`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	defer log.Flush()
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugsorigins")
	result :=  make([]interface{}, 0)
	err := collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Info(err)
	}
	log.Info(result)
	return result
}

func GetDrugMarketingDates() []interface{} {
	defer log.Flush()
	session := connectToDb()
	defer session.Close()
	collection := session.DB("medigo").C("drugsdates")
	result :=  make([]interface{}, 0)
	err := collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Info(err)
	}
	return result
}
