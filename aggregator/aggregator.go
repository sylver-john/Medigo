package aggregator

import (
	m "github.com/sylver-john/Medigo/src/model"
	"strings"
)

func GetDrugOriginPercent(drugOrigins []m.Titulaire) map[string]float32 {
	percent := make(map[string]float32)
	for _, value := range drugOrigins {
		for _, titulaire := range value.Titulaire {
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
	for _, date := range drugDates {
			s := strings.Split(date.DateMiseSurLeMarche, "/")
			if len(s) > 2 {
				if  _, ok := percent[s[2]]; ok {
					percent[s[2]] += 1
				} else {
					percent[s[2]] = 0
					percent[s[2]] += 1
				}
			}
		}
	return percent
}
