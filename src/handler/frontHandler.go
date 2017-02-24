package handler

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.File("C:/Users/simhoff/Medigo/front/index.html")
}
