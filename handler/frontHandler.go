package handler

import (
	"github.com/gin-gonic/gin"
	//"log"
)

func Index(c *gin.Context) {
	c.File("C:/Projects/src/Medigo/front/index.html")
}
