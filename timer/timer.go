package timer

import (
  "fmt"

  mod "github.com/OOM-Killer/fakemetrics_ng/timer/module"
  rt "github.com/OOM-Killer/fakemetrics_ng/timer/realtime"
  bf "github.com/OOM-Killer/fakemetrics_ng/timer/backfill"
)

var(
  moduleMap []*mod.ModuleT = []*mod.ModuleT{
    rt.Module,
    bf.Module,
  }
)

func RegisterFlagSets() {
  for _,t := range moduleMap {
    t.RegFlags()
  }
}

func GetInstance(seek string, agentId int) (mod.Timer) {
  for _,t := range moduleMap {
    if t.Name == seek {
      return t.Init(agentId)
    }
  }
  panic(fmt.Sprintf("failed to find timer %s", seek))
}
