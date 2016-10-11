package out

import (
  "fmt"

  "github.com/OOM-Killer/fakemetrics_ng/out/iface"
  mp "github.com/OOM-Killer/fakemetrics_ng/out/multiplexer"
  carbon "github.com/OOM-Killer/fakemetrics_ng/out/carbon"
)

var (
  modules = []iface.OutIface{
    &carbon.Carbon{},
  }
)

func RegisterFlagSets() {
  for _,o := range modules {
    o.RegisterFlagSet()
  }
}

func GetInstance(name string) (iface.OutIface) {
  for _,o := range modules {
    if o.GetName() == name {
      return o
    }
  }
  panic(fmt.Sprintf("failed to find output %s", name))
}

func GetMultiInstance(names []string) (iface.OutIface) {
  m := mp.Multiplexer{}
  for _,name := range names {
    m.AddOut(GetInstance(name))
  }
  return &m
}
