package handler

import (
	"log"
	"github.com/gin-gonic/gin"
)


func Index(c *gin.Context) {
	log.Println("Serving the file")
  	c.File("C:/Users/simhoff/Medigo/front/index.html")
}
