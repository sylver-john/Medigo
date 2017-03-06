package service

import(
  log "github.com/cihub/seelog"
  "io/ioutil"
  "encoding/json"
  m "github.com/sylver-john/Medigo/src/model"
)

func GetConf() m.Config {
  defer log.Flush()
  var config m.Config
	rawConfig, err := ioutil.ReadFile("C:/users/simhoff/Perso/src/Medigo/conf.json")
	if err != nil {
		log.Info(err)
	}
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
    log.Info(err)
	}
	return config
}
