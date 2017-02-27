package service

import (
	m "Medigo/src/model"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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

func GetDrugByCIS(cis string) m.Drug {
	resp, err := http.Get("https://medicaments.api.gouv.fr/api/medicaments/" + cis)
	log.Println("https://medicaments.api.gouv.fr/api/medicaments/" + cis)
	var drug m.Drug
	if err != nil && resp.StatusCode != 200 {
		log.Println(err)
		return drug
	}
	err = json.NewDecoder(resp.Body).Decode(&drug)
	if err != nil {
		log.Println(err)
	}
	return drug
}

func GetDrugs(cis []string) []m.Drug {
	var drugs []m.Drug
	for i := 0; i < len(cis); i++ {
		drugs = append(drugs, GetDrugByCIS(cis[i]))
	}
	return drugs
}

func GetDrugsAndStore(cisList <-chan string, results chan<- bool) {
	for cis := range cisList {
		drugs := GetDrugByCIS(cis)
		StoreDrug(drugs)
		results <- false
	}
}
