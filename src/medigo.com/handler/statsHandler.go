package handler

import (
	"medigo.com/service"
	"github.com/gin-gonic/gin"
)

func DrugsOrigins(c *gin.Context) {
	drugsOriginsPercent := service.GetDrugsOrigins()
	c.JSON(200, drugsOriginsPercent)
}

func DrugsDates(c *gin.Context) {
	drugsDatesPercent := service.GetDrugMarketingDates()
	c.JSON(200, drugsDatesPercent)
}
