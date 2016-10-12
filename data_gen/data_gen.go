package data_gen

import (
  "fmt"

  "gopkg.in/raintank/schema.v1"

  simple "github.com/OOM-Killer/fakemetrics_ng/data_gen/simple"
)

type DataGen interface {
  GetData(int64) ([]*schema.MetricData)
}

type Module struct {
  Name      string
  Init      func() (DataGen)
  RegFlags  func()
}

var(
  moduleMap []Module = []Module{
    {
      "simple",
      func() (DataGen) {return &simple.Simple{}},
      simple.RegisterFlagSet,
    },
  }
)

func RegisterFlagSets() {
  for _,dg := range moduleMap {
    dg.RegFlags()
  }
}

func GetInstance(seek string) (DataGen) {
  for _,dg := range moduleMap {
    if dg.Name == seek {
      return dg.Init()
    }
  }
  panic(fmt.Sprintf("failed to find data_gen %s", seek))
}
