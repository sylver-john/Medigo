package handler

import (
	"../aggregator"
	"../service"
	"github.com/gin-gonic/gin"
	"log"
)

func Index(c *gin.Context) {
	inventory := service.GetInventory()
	drugs := service.GetDrugs(inventory[0:200])
	for _, drugCollection := range drugs {
		if drugCollection != nil && 0 < len(drugCollection) {
			log.Println(drugCollection)
			service.StoreDrugCollection(drugCollection)
		}
	}
	c.File("C:/Users/simhoff/Medigo/front/index.html")
}

func DrugsOrigins(c *gin.Context) {
	drugsOrigins := service.GetDrugsorigins()
	drugsOriginsPercent := aggregator.GetDrugOriginPercent(drugsOrigins)
	c.JSON(200, drugsOriginsPercent)
}
