package timer

import (
  "fmt"

  mod "github.com/OOM-Killer/fakemetrics_ng/timer/module"
  rt "github.com/OOM-Killer/fakemetrics_ng/timer/realtime"
)

var(
  moduleMap []*mod.ModuleT = []*mod.ModuleT{
    rt.Module,
  }
)

func RegisterFlagSets() {
  for _,t := range moduleMap {
    t.RegFlags()
  }
}

func GetInstance(seek string) (mod.Timer) {
  for _,t := range moduleMap {
    if t.Name == seek {
      return t.Init()
    }
  }
  panic(fmt.Sprintf("failed to find timer %s", seek))
}
