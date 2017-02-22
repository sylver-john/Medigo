package handler

import (
	"log"
	"github.com/gin-gonic/gin"
	"../service"
)


func Index(c *gin.Context) {
	log.Println("index")
	inventory := service.GetInventory()
	drugs := service.GetDrugs(inventory[0:2])
	for _, value := range drugs {
		if value != nil {
			log.Println(value[0].Nom)
		}
	}
  	c.File("C:/Users/simhoff/Medigo/front/index.html")
}
