package handler

import (
	"log"
	"github.com/gin-gonic/gin"
	"../service"
)


func Index(c *gin.Context) {
	inventory := service.GetInventory()
	drugs := service.GetDrugs(inventory[0:200])
	for _, value := range drugs {
		if value != nil && 0 < len(value) {
			log.Println(value)
			service.StoreDrug(value)
		}
	}
  	c.File("C:/Users/simhoff/Medigo/front/index.html")
}
