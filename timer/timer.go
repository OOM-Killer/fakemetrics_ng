package timer

import (
  "time"
  "fmt"

  rt "github.com/OOM-Killer/fakemetrics_ng/timer/realtime"
)

var (
  modules = []Timer{
    &rt.Realtime{},
  }
)

type Timer interface {
  RegisterFlagSet()
  GetName() (string)
  GetTicker() (*time.Ticker)
  GetTimestamp() (int64)
}

func RegisterFlagSets() {
  for _,t := range modules {
    t.RegisterFlagSet()
  }
}

func GetInstance(name string) (Timer) {
  for _,t := range modules {
    if t.GetName() == name {
      return t
    }
  }
  panic(fmt.Sprintf("failed to find timer %s", name))
}
