package main

import (
  "fmt"
  "flag"
  //"os"
  //"time"

  "github.com/OOM-Killer/fakemetrics_ng/timer"
  "github.com/OOM-Killer/fakemetrics_ng/data_gen"
)

var (
  confFile = flag.String("config",
                         "fakemetrics.ini",
                         "configuration file path")
)

func main() {
  flag.Parse()
  timerFactory := timer.NewFactory()
  dataGenFactory := data_gen.NewFactory()

  setupConfig()

  timer := timerFactory.GetTimer(timerMod)
  dataGen := dataGenFactory.GetDataGen(dataGenMod)

  tick := timer.GetTicker()
  for range tick.C {
    fmt.Println(dataGen.GetData)
  }

}
