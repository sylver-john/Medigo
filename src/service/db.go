package service

import(
  "log"
  m "../model"
  "gopkg.in/mgo.v2"
)

func connectToDb() *mgo.Session {
  session, err := mgo.Dial("**********")
  if err != nil {
    log.Fatal(err)
  }
  return session
}

func StoreDrug(drugs []m.Drug) {
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
