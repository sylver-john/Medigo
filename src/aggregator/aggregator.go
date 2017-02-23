package aggregator

import (
	m "../model"
	"log"
	"strings"
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
	percent["2000"] = 0
	percent["2001"] = 0
	percent["2002"] = 0
	percent["2003"] = 0
	percent["2004"] = 0
	percent["2005"] = 0
	percent["2006"] = 0
	percent["2007"] = 0
	percent["2008"] = 0
	percent["2009"] = 0
	percent["2010"] = 0
	percent["2011"] = 0
	percent["2012"] = 0
	percent["2013"] = 0
	percent["2014"] = 0
	percent["2015"] = 0
	percent["2016"] = 0
	percent["2017"] = 0
	for _, date := range drugDates {
		s := strings.Split(date.DateMiseSurLeMarche, "/")
		log.Println(s)
		log.Println(s[0])
		if _, ok := percent[s[2]]; ok {
			percent[s[2]] += 1
		}
	}
	return percent
}
