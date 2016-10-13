package module

import (
  "gopkg.in/raintank/schema.v1"
)

type DataGen interface {
  GetData(int64) ([]*schema.MetricData)
}

type ModuleT struct {
  Name      string
  Init      func() (DataGen)
  RegFlags  func()
}
