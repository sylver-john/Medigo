package service

import(
	"encoding/json"
	"io/ioutil"
	log "github.com/cihub/seelog"
)
// Récupérer la liste des noms

func GetInventory() []string {
	defer log.Flush()
	content, err := ioutil.ReadFile("C:/Projects/src/Medigo/cis.json")
	if err != nil {
		log.Info(err)
	}
	var inventory []string
	err = json.Unmarshal(content, &inventory)
	if err != nil {
		log.Info(err)
	}
	return inventory
}
