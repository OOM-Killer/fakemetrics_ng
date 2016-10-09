package main

import (
  "fmt"
  "flag"

  "gopkg.in/raintank/schema.v1"

  "github.com/OOM-Killer/fakemetrics_ng/timer"
  "github.com/OOM-Killer/fakemetrics_ng/data_gen"
  "github.com/OOM-Killer/fakemetrics_ng/out"
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
  outFactory := out.New()

  setupConfig()

  timer := timerFactory.GetInstance(timerMod)
  dataGen := dataGenFactory.GetInstance(dataGenMod)
  out := outFactory.GetInstance(outMod)

  outChan := out.GetChan()
  tick := timer.GetTicker()
  for range tick.C {
    go doTick(&dataGen, outChan, timer.GetTimestamp())
  }
}

func doTick(dg *data_gen.DataGen, outChan chan *schema.MetricData, ts int64) {
  metric := (*dg).GetData(ts)
  fmt.Println(metric)
  outChan<-metric
}
