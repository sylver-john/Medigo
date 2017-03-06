package handler

import (
	"github.com/sylver-john/Medigo/src/aggregator"
	"github.com/sylver-john/Medigo/src/service"
	"github.com/gin-gonic/gin"
)

func DrugsOriginsFromBase(c *gin.Context) {
	drugsOrigins := service.GetDrugsOriginsFromBase()
	drugsOriginsPercent := aggregator.GetDrugOriginPercent(drugsOrigins)
	service.StoreDrugsOrigins(drugsOriginsPercent)
	c.JSON(200, drugsOriginsPercent)
}

func DrugsDatesFromBase(c *gin.Context) {
	drugsDates := service.GetDrugMarketingDatesFromBase()
	drugsDatesPercent := aggregator.GetdrugsDatesPercent(drugsDates)
	service.StoreDrugsDates(drugsDatesPercent)
	c.JSON(200, drugsDatesPercent)
}

func BaseUpdateGoroutine(c *gin.Context) {
	names := make(chan string, 100000)
	results := make(chan bool, 100000)

	inventory := service.GetInventory()
	for i := 0; i < 10; i++ {
		go service.GetDrugsAndStore(names, results)
	}
	for j := 0; j < len(inventory); j++ {
		names <- inventory[j]
	}
	close(names)
	for k := 0; k < len(inventory); k++ {
		<-results
	}
	close(results)
}
