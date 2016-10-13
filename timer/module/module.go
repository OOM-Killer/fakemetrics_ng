package module

import (
  "time"
)

type Timer interface {
  GetInterval() (int)
  GetTicker() (*time.Ticker)
  GetTimestamp() (int64)
}

type ModuleT struct {
  Name      string
  Init      func(int) (Timer)
  RegFlags  func()
}
