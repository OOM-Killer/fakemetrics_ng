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
