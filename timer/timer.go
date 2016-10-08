package timer

import (
  "time"
  "fmt"

  rt "github.com/OOM-Killer/fakemetrics_ng/timer/realtime"
)

type Timer interface {
  PrintInterval()
  RegisterFlagSet()
  GetTicker() (*time.Ticker)
  GetName() (string)
}

type TimerFactory struct {
  timers []Timer
}

func NewFactory() TimerFactory {
  inst := TimerFactory{}
  inst.initTimers()
  inst.registerFlagSets()
  return inst
}

func (f *TimerFactory) initTimers() {
  f.timers = []Timer{
    &rt.Realtime{},
  }
}

func (f *TimerFactory) registerFlagSets() {
  for _, timer := range f.timers {
    timer.RegisterFlagSet()
  }
}

func (f *TimerFactory) GetTimer(seek string) (Timer) {
  f.initTimers()
  for _,timer := range f.timers {
    if (timer.GetName() == seek) {
      return timer
    }
  }
  panic(fmt.Sprintf("could not find timer %s", seek))
}
