package iface

import (
  "gopkg.in/raintank/schema.v1"
)

type OutIface interface {
  Start()
  Put(*schema.MetricData)
}
