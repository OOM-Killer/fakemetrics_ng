package data_gen

import (
  "fmt"

  mod "github.com/OOM-Killer/fakemetrics_ng/data_gen/module"
  simple "github.com/OOM-Killer/fakemetrics_ng/data_gen/simple"
)

var(
  moduleMap []*mod.ModuleT = []*mod.ModuleT{
    simple.Module,
  }
)

func RegisterFlagSets() {
  for _,dg := range moduleMap {
    dg.RegFlags()
  }
}

func GetInstance(seek string) (mod.DataGen) {
  for _,dg := range moduleMap {
    if dg.Name == seek {
      return dg.Init()
    }
  }
  panic(fmt.Sprintf("failed to find data_gen %s", seek))
}
