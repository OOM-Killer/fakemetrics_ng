package main

import (
  "fmt"
  "flag"

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
  timerFactory := timer.New()
  dataGenFactory := data_gen.New()

  setupConfig()

  fmt.Println("getting data generator " + dataGenMod)
  timer := timerFactory.GetInstance(timerMod)
  dataGen := dataGenFactory.GetInstance(dataGenMod)

  tick := timer.GetTicker()
  for range tick.C {
    fmt.Println(dataGen.GetData())
  }

}
