package handler

import (
	"log"
	"github.com/gin-gonic/gin"
	"../service"
)


func Index(c *gin.Context) {
	inventory := service.GetInventory()
	drugs := service.GetDrugs(inventory[0:20])
	for _, value := range drugs {
		if value != nil {
			log.Println(value)
		}
	}
  	c.File("C:/Users/simhoff/Medigo/front/index.html")
}
