package timer

import (
  "time"
  "fmt"

  rt "github.com/OOM-Killer/fakemetrics_ng/timer/realtime"
)

type Timer interface {
  GetInterval() (int)
  GetTicker() (*time.Ticker)
  GetTimestamp() (int64)
}

type Module struct {
  Name      string
  Init      func() (Timer)
  RegFlags  func()
}

var(
  moduleMap []Module = []Module{
    {
      "realtime",
      func() (Timer) {return &rt.Realtime{}},
      rt.RegisterFlagSet,
    },
  }
)

func RegisterFlagSets() {
  for _,t := range moduleMap {
    t.RegFlags()
  }
}

func GetInstance(seek string) (Timer) {
  for _,t := range moduleMap {
    if t.Name == seek {
      return t.Init()
    }
  }
  panic(fmt.Sprintf("failed to find timer %s", seek))
}
