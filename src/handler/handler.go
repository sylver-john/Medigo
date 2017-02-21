package handler

import (
	"log"
	"github.com/gin-gonic/gin"
	"../service"
)


func Index(c *gin.Context) {
	log.Println("index")
	service.GetInventory()
  	c.File("C:/Users/simhoff/Medigo/front/index.html")
}
