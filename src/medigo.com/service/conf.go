package service

import(
  "log"
  "io/ioutil"
  "encoding/json"
  m "medigo.com/model"
)

func GetConf() m.Config {
  var config m.Config
	rawConfig, err := ioutil.ReadFile("C:/Users/sylver/Medigo/conf.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
    log.Fatal(err)
	}
	return config
}
