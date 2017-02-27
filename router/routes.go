package router

import (
  "github.com/gin-gonic/gin"
	"Medigo/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/index",
		handler.Index,
	},
  Route{
    "GetDrugsOrigins",
    "GET",
    "/drugsorigins",
    handler.GetDrugsOrigins,
  },
  Route{
    "GetDrugsDates",
    "GET",
    "/drugsdates",
    handler.GetDrugsDates,
  },
}

/*
Route{
  "DrugsDatesFromBase",
  "GET",
  "/drugsdatesfrombase",
  handler.DrugsDatesFromBase,
},
Route{
  "DrugsOriginsFromBase",
  "GET",
  "/drugsoriginsfrombase",
  handler.DrugsOriginsFromBase,
},
Route{
  "BaseUpdateGoroutine",
  "GET",
  "/baseupdategoroutine",
  handler.BaseUpdateGoroutine,
},
*/
