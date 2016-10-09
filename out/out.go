package out

import (
  iface "github.com/OOM-Killer/fakemetrics_ng/out/interface"
  fact "github.com/OOM-Killer/fakemetrics_ng/factory"
  mp "github.com/OOM-Killer/fakemetrics_ng/out/multiplexer"
)

var (
  modules = []iface.OutIface{
  }
)

type OutFactory struct {
  fact.Factory
}

func New() (OutFactory) {
  fact := OutFactory{}
  for _,mod := range modules {
    fact.Factory.RegisterModule(mod)
  }

  fact.Factory.RegisterFlagSets()
  return fact
}

func (f *OutFactory) GetSingleInstance (name string) (*iface.OutIface) {
  inst := f.Factory.GetInstance(name).(iface.OutIface)
  return &inst
}

func (f *OutFactory) GetInstance(names []string) (iface.OutIface) {
  m := mp.Multiplexer{}
  for _,name := range names {
    m.AddOut(f.GetSingleInstance(name))
  }
  return &m
}
