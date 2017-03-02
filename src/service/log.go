package service

import 	log "github.com/cihub/seelog"

func SetLogger() {
  testConfig := `
  <seelog type="asyncloop">
  <outputs>
    <filter levels="trace,debug">
      <file path="./log/debug.log"/>
    </filter>
    <filter levels="info">
      <file path="./log/info.log"/>
    </filter>
    <filter levels="warn">
      <file path="./log/warn.log"/>
    </filter>
    <filter levels="error,critical">
      <file path="./log/error.log"/>
    </filter>
  </outputs>
  </seelog>
  `
  logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
  log.ReplaceLogger(logger)
}
