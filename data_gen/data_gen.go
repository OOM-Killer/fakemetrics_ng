package data_gen

import (
  "gopkg.in/raintank/schema.v1"

  fact "github.com/OOM-Killer/fakemetrics_ng/factory"
  simple "github.com/OOM-Killer/fakemetrics_ng/data_gen/simple"
)

var (
  modules = []DataGen{
    &simple.Simple{},
  }
)

type DataGen interface {
  RegisterFlagSet()
  GetData() (*schema.MetricData)
  GetName() (string)
}

type DataGenFactory struct {
  fact.Factory
}

func New() DataGenFactory {
  fact := DataGenFactory{}
  for _,mod := range modules {
    fact.Factory.RegisterModule(mod)
  }

  fact.Factory.RegisterFlagSets()
  return fact
}

func (f *DataGenFactory) GetInstance (name string) (DataGen) {
  return f.Factory.GetInstance(name).(DataGen)
}
