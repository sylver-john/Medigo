package service

import (
	m "../model"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"unicode/utf8"
)

func GetDrugByName(name string) []m.Drug {
	Url, err := url.Parse("https://medicaments.api.gouv.fr/api/medicaments")
	if err != nil {
		log.Println(err)
	}
	params := url.Values{}
	params.Add("nom", name)
	Url.RawQuery = params.Encode()
	log.Println(Url.String())
	resp, err := http.Get(Url.String())
	if err != nil && resp.StatusCode != 200 {
		log.Println(err)
		return nil
	}
	log.Print(resp.StatusCode)
	var drug []m.Drug
	err = json.NewDecoder(resp.Body).Decode(&drug)
	if err != nil {
		log.Println(err)
	}
	return drug
}

func GetDrugByCIS(cis string) {
	resp, err := http.Get("https://medicaments.api.gouv.fr/api/medicaments/" + cis)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

func GetDrugs(names []string) m.Drugs {
	var drugs m.Drugs
	for i := 0; i < len(names); i++ {
		if utf8.ValidString(names[i]) {
			drugs = append(drugs, GetDrugByName(names[i]))
		} else {
			log.Println("invalid format name")
		}
	}
	return drugs
}

func GetDrugsAndStore(names <-chan string, results chan<- bool) {
	for name := range names {
		drugCollection := GetDrugByName(name)
		if drugCollection != nil && 0 < len(drugCollection) {
			StoreDrugCollection(drugCollection)
			results<- true
		} else {
			results<- false
		}
	}
}
