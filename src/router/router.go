package router

import (
  	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
  router.Use(gin.Logger())
  router.Use(gin.Recovery())
	for _, route := range routes {
		var handler gin.HandlerFunc

		handler = route.HandlerFunc
		router.Handle(route.Method, route.Pattern, handler)
	}
	return router
}
