package handler

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
)

func Index(c *gin.Context) {
	dir, _ := os.Getwd()
	log.Println(dir)
	c.File(dir + "/front/index.html")
}
