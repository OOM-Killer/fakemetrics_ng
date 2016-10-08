package main

import (
  "fmt"
  "flag"
  //"os"
  //"time"

  "github.com/OOM-Killer/fakemetrics_ng/timer"
)

var (
  confFile = flag.String("config",
                         "fakemetrics.ini",
                         "configuration file path")
)

func main() {
  flag.Parse()
  timerFactory := timer.NewFactory()

  setupConfig()

  timer := timerFactory.GetTimer(timerMod)
  timer.PrintInterval()

  tick := timer.GetTicker()
  for range tick.C {
    fmt.Println("tick")
  }

}
