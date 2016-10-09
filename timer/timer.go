package timer

import (
  "time"

  fact "github.com/OOM-Killer/fakemetrics_ng/factory"
  rt "github.com/OOM-Killer/fakemetrics_ng/timer/realtime"
)

var (
  modules = []Timer{
    &rt.Realtime{},
  }
)

type Timer interface {
  fact.Module
  GetTicker() (*time.Ticker)
  GetTimestamp() (int64)
}

type TimerFactory struct {
  fact.Factory
}

func New() (TimerFactory) {
  fact := TimerFactory{}
  for _,mod := range modules {
    fact.Factory.RegisterModule(mod)
  }

  fact.Factory.RegisterFlagSets()
  return fact
}

func (f *TimerFactory) GetInstance (name string) (Timer) {
  return f.Factory.GetInstance(name).(Timer)
}
