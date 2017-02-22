package service

import (
	"log"
	"net/http"
	"net/url"
	"encoding/json"
	m "../model"
)

func GetDrugByName(name string) []m.Drug {
	url, err := url.Parse("https://medicaments.api.gouv.fr/api/medicaments?nom=" + name)
	if err != nil {
		log.Println(err)
	}
	log.Println(url.String())
	resp, err := http.Get(url.String())
	if err != nil && resp.StatusCode != 200 {
		log.Println(err)
		return nil
	}
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
	for i := 0; i <len(names); i++ {
		drugs = append(drugs, GetDrugByName(names[i]))
	}
	return drugs
}
