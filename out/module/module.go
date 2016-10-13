package module

import (
  "gopkg.in/raintank/schema.v1"
)

type OutIface interface {
  Start()
  Put(*schema.MetricData)
}

type ModuleT struct {
  Name      string
  Init      func() (OutIface)
  RegFlags  func()
}
