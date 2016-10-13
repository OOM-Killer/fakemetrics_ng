package agents

import (
  "time"

  timer "github.com/OOM-Killer/fakemetrics_ng/timer/module"
  data_gen "github.com/OOM-Killer/fakemetrics_ng/data_gen/module"
  out "github.com/OOM-Killer/fakemetrics_ng/out/module"
)

type Agent struct {
  timer timer.Timer
  dataGen data_gen.DataGen
  out out.OutIface
  offset int
}

func (a *Agent) Run() {
  time.Sleep(time.Duration(a.offset))

  a.out.Start()
  tick := a.timer.GetTicker()
  for range tick {
    go a.doTick()
  }
}

func (a *Agent) doTick() {
  metrics := a.dataGen.GetData(a.timer.GetTimestamp())
  for _,m := range metrics {
    a.out.Put(m)
  }
}
