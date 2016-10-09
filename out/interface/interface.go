package OutIface

import (
  fact "github.com/OOM-Killer/fakemetrics_ng/factory"

  "gopkg.in/raintank/schema.v1"
)

type OutIface interface {
  fact.Module
  GetChan() (chan *schema.MetricData)
}
