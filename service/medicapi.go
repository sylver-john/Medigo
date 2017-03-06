package service

import (
	m "github.com/sylver-john/Medigo/model"
	"encoding/json"
	log "github.com/cihub/seelog"
	"net/http"
	"net/url"
)

func GetDrugByName(name string) []m.Drug {
	SetLogger()
	defer log.Flush()
	Url, err := url.Parse("https://medicaments.api.gouv.fr/api/medicaments")
	if err != nil {
		log.Warn(err)
	}
	params := url.Values{}
	params.Add("nom", name)
	Url.RawQuery = params.Encode()
	log.Info(Url.String())
	resp, err := http.Get(Url.String())
	if err != nil && resp.StatusCode != 200 {
		log.Warn(err)
		return nil
	}
	log.Info(resp.StatusCode)
	var drug []m.Drug
	err = json.NewDecoder(resp.Body).Decode(&drug)
	if err != nil {
		log.Warn(err)
	}
	return drug
}

func GetDrugByCIS(cis string) m.Drug {
	SetLogger()
	defer log.Flush()
	resp, err := http.Get("https://medicaments.api.gouv.fr/api/medicaments/" + cis)
	log.Info("https://medicaments.api.gouv.fr/api/medicaments/" + cis)
	var drug m.Drug
	if err != nil && resp.StatusCode != 200 {
		log.Warn(err)
		return drug
	}
	err = json.NewDecoder(resp.Body).Decode(&drug)
	if err != nil {
		log.Warn(err)
	}
	return drug
}

func GetDrugs(cis []string) []m.Drug {
	SetLogger()
	var drugs []m.Drug
	for i := 0; i < len(cis); i++ {
		drugs = append(drugs, GetDrugByCIS(cis[i]))
	}
	return drugs
}

func GetDrugsAndStore(cisList <-chan string, results chan<- bool) {
	SetLogger()
	for cis := range cisList {
		drugs := GetDrugByCIS(cis)
		StoreDrug(drugs)
		results <- false
	}
}
