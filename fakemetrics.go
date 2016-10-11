package main

import (
  "flag"


  "github.com/OOM-Killer/fakemetrics_ng/agents"
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
  timer.RegisterFlagSets()
  data_gen.RegisterFlagSets()
  out.RegisterFlagSets()
  agents.RegisterFlagSets()

  setupConfig()

  a := agents.New(timerMod, dataGenMod, outMod)
  a.Run()
  /*timer := timer.GetInstance(timerMod)
  dataGen := data_gen.GetInstance(dataGenMod)
  out := out.GetMultiInstance(outMod)

  out.Start()

  tick := timer.GetTicker()
  for range tick.C {
    go doTick(dataGen, out, timer.GetTimestamp())
  }*/
}

/*func doTick(dg data_gen.DataGen, out iface.OutIface, ts int64) {
  metrics := dg.GetData(ts)
  for _,m := range metrics {
    out.Put(m)
  }
}*/
