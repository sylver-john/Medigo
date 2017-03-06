package service

import(
	"encoding/json"
	"io/ioutil"
	log "github.com/cihub/seelog"
)
// Récupérer la liste des noms

func GetInventory() []string {
	SetLogger()
	defer log.Flush()
	content, err := ioutil.ReadFile("C:/Projects/src/Medigo/cis.json")
	if err != nil {
		log.Warn(err)
	}
	var inventory []string
	err = json.Unmarshal(content, &inventory)
	if err != nil {
		log.Warn(err)
	}
	return inventory
}
