package aggregator

import (
	"log"
	//"time"
	m "../model"
)

func GetDrugOriginPercent(drugOrigins []m.Titulaire) map[string]float32 {
	percent := make(map[string]float32)
	for _, value := range drugOrigins {
		log.Println(value)
		for _, titulaire := range value.Titulaire {
			log.Println(titulaire)
			if _, ok := percent[titulaire]; ok {
				percent[titulaire] += 1
			} else {
				percent[titulaire] = 0
				percent[titulaire] += 1
			}
		}
	}
	return percent
}

func GetdrugsDatesPercent(drugDates []m.DateMiseSurLeMarche) map[string]float32 {
		percent := make(map[string]float32)
		percent["2000-2005"] = 100
		return percent
}
