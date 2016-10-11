package data_gen

import (
  "fmt"

  "gopkg.in/raintank/schema.v1"

  simple "github.com/OOM-Killer/fakemetrics_ng/data_gen/simple"
)

var (
  modules = []DataGen{
    &simple.Simple{},
  }
)

type DataGen interface {
  RegisterFlagSet()
  GetName() (string)
  GetData(int64) ([]*schema.MetricData)
}

func RegisterFlagSets() {
  for _,dg := range modules {
    dg.RegisterFlagSet()
  }
}

func GetInstance(name string) (DataGen) {
  for _,dg := range modules {
    if dg.GetName() == name {
      return dg
    }
  }
  panic(fmt.Sprintf("failed to find data_gen %s", name))
}
