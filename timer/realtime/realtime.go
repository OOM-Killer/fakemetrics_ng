package realtime

import (
  "flag"
  "time"

  mod "github.com/OOM-Killer/fakemetrics_ng/timer/module"
  gc "github.com/rakyll/globalconf"
)

var (
  interval int
)

var Module *mod.ModuleT = &mod.ModuleT{
  "realtime",
  func(id int) (mod.Timer) {return &Realtime{id}},
  RegisterFlagSet,
}

type Realtime struct {
  agentId int
}

func RegisterFlagSet() {
  flags := flag.NewFlagSet("realtime", flag.ExitOnError)
  flags.IntVar(&interval, "interval", 100, "the metric interval")
  gc.Register("realtime", flags)
}

func (r *Realtime) GetTicker() (*time.Ticker) {
  return time.NewTicker(time.Duration(interval) * time.Millisecond)
}

func (r *Realtime) GetTimestamp() (int64) {
  return time.Now().Unix()
}

func (r *Realtime) GetInterval() (int) {
  return interval
}
