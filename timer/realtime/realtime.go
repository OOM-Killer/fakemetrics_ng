package realtime

import (
  "flag"
  "time"

  gc "github.com/rakyll/globalconf"
)

type Realtime struct {}

var (
  interval int
)

func (r *Realtime) RegisterFlagSet() {
  rtFlags := flag.NewFlagSet(r.GetName(), flag.ExitOnError)
  rtFlags.IntVar(&interval, "interval", 100, "the metric interval")
  gc.Register(r.GetName(), rtFlags)
}

func (r *Realtime) GetTicker() (*time.Ticker) {
  return time.NewTicker(time.Duration(interval) * time.Millisecond)
}

func (r *Realtime) GetName() (string) {
  return "realtime"
}
