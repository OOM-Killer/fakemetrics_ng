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
  flags := flag.NewFlagSet(r.GetName(), flag.ExitOnError)
  flags.IntVar(&interval, "interval", 100, "the metric interval")
  gc.Register(r.GetName(), flags)
}

func (r *Realtime) GetName() (string) {
  return "realtime"
}

func (r *Realtime) GetTicker() (*time.Ticker) {
  return time.NewTicker(time.Duration(interval) * time.Millisecond)
}
