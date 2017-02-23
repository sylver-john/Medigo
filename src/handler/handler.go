package handler

import (
	"../aggregator"
	"../service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Index(c *gin.Context) {
	c.File("C:/Users/simhoff/Medigo/front/index.html")
}

func DrugsOrigins(c *gin.Context) {
	drugsOrigins := service.GetDrugsorigins()
	drugsOriginsPercent := aggregator.GetDrugOriginPercent(drugsOrigins)
	c.JSON(200, drugsOriginsPercent)
}

func BaseUpdate(c *gin.Context) {
	inventory := service.GetInventory()
	drugs := service.GetDrugs(inventory[0:200])
	for _, drugCollection := range drugs {
		if drugCollection != nil && 0 < len(drugCollection) {
			service.StoreDrugCollection(drugCollection)
		} else {
			log.Println("drug is nil")
		}
	}
}

func BaseUpdateGoroutine(c *gin.Context) {
	log.Println(time.Now())
	names := make(chan string, 100000)
	results := make(chan bool, 100000)

	inventory := service.GetInventory()
	for i := 0; i < 10; i++ {
		go service.GetDrugsAndStore(names, results)
	}
	for j := 0; j < 200; j++ {
		names<- inventory[j]
	}
	close(names)
	for k := 0; k < 200; k++ {
		<-results
	}
	close(results)
	log.Println(time.Now())
}
