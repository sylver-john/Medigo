package service

import (
	"log"
	"net/http"
)

func GetDrugByName(name string) {
	resp, err := http.Get("https://medicaments.api.gouv.fr/api/medicaments?nom=" + name)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

func GetDrugByCIS(cis string) {
	resp, err := http.Get("https://medicaments.api.gouv.fr/api/medicaments/" + cis)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}