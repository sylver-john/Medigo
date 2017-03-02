package handler

import (
	"Medigo/src/service"
	"github.com/gin-gonic/gin"
)

func GetDrugsOrigins(c *gin.Context) {
	drugsOriginsPercent := service.GetDrugsOrigins()
	c.JSON(200, drugsOriginsPercent)
}

func GetDrugsDates(c *gin.Context) {
	drugsDatesPercent := service.GetDrugMarketingDates()
	c.JSON(200, drugsDatesPercent)
}