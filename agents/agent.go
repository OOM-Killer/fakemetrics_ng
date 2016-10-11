package agents

import (
  "time"

  "github.com/OOM-Killer/fakemetrics_ng/timer"
  "github.com/OOM-Killer/fakemetrics_ng/data_gen"
  "github.com/OOM-Killer/fakemetrics_ng/out/iface"
)

type Agent struct {
  timer timer.Timer
  dataGen data_gen.DataGen
  out iface.OutIface
  offset int
}

func (a *Agent) Run() {
  time.Sleep(time.Duration(a.offset))

  tick := a.timer.GetTicker()
  a.out.Start()
  for range tick.C {
    go a.doTick()
  }
}

func (a *Agent) doTick() {
  metrics := a.dataGen.GetData(a.timer.GetTimestamp())
  for _,m := range metrics {
    a.out.Put(m)
  }
}
