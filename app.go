package main

import (
  "Medigo/src/router"
	log "github.com/cihub/seelog"
  "Medigo/src/service"
)

func main() {
  service.SetLogger()
  defer log.Flush()
  log.Info("Starting the server")
  router := router.NewRouter()
  router.Run()
}
