package service

import(
	//"../model"
	"encoding/json"
	"io/ioutil"
	"log"
)
// Récupérer la liste des noms

func GetInventory() []string {
	content, err := ioutil.ReadFile("C:/Users/simhoff/cis.json")
	if err != nil {
		log.Println(err)
	}
	var inventory []string
	err = json.Unmarshal(content, &inventory)
	if err != nil {
		log.Println(err)
	}
	return inventory
}
