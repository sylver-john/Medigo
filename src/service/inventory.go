package service

import(
	"../model"
	"encoding/json"
	"io/ioutil"
	"log"
)
// Récupérer la liste des noms

func GetInventory() {
	content, err := ioutil.ReadFile("C:/Users/simhoff/COMPO.json")
	if err != nil {
		log.Println(err)
	}
	var inventory model.Drugs
	err = json.Unmarshal(content, &inventory)
	if err != nil {
		log.Println(err)
	}
	log.Println(inventory)
}