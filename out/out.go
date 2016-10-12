package out

import (
  "fmt"

  "github.com/OOM-Killer/fakemetrics_ng/out/iface"
  mp "github.com/OOM-Killer/fakemetrics_ng/out/multiplexer"
  carbon "github.com/OOM-Killer/fakemetrics_ng/out/carbon"
)

type Module struct {
  Name      string
  Init      func() (iface.OutIface)
  RegFlags  func()
}

var (
  moduleMap []Module = []Module{
    {
      "carbon",
      func() (iface.OutIface) {return &carbon.Carbon{}},
      carbon.RegisterFlagSet,
    },
  }
)

func RegisterFlagSets() {
  for _,o := range moduleMap {
    o.RegFlags()
  }
}

func GetInstance(seek string) (iface.OutIface) {
  for _,o := range moduleMap {
    if o.Name == seek {
      return o.Init()
    }
  }
  panic(fmt.Sprintf("failed to find output %s", seek))
}

func GetMultiInstance(names []string) (iface.OutIface) {
  m := mp.Multiplexer{}
  for _,name := range names {
    m.AddOut(GetInstance(name))
  }
  return &m
}
