package iface

import (
  "gopkg.in/raintank/schema.v1"
)

type OutIface interface {
  RegisterFlagSet()
  GetName() (string)
  Start()
  Put(*schema.MetricData)
}
