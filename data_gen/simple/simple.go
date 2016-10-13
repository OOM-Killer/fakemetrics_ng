package simple

import (
  "flag"
  "fmt"

  "gopkg.in/raintank/schema.v1"

  mod "github.com/OOM-Killer/fakemetrics_ng/data_gen/module"
  gc "github.com/rakyll/globalconf"
)

var (
  keyCount int
  keyPrefix string
)

var Module *mod.ModuleT = &mod.ModuleT{
  "simple",
  func() (mod.DataGen) {return &Simple{}},
  RegisterFlagSet,
}

type Simple struct {}

func RegisterFlagSet() {
  flags := flag.NewFlagSet("simple", flag.ExitOnError)
  flags.IntVar(&keyCount, "key_count", 100, "number of keys to generate")
  flags.StringVar(&keyPrefix, "key_prefix", "some.key.", "prefix for keys")
  gc.Register("simple", flags)
}

func (s *Simple) GetData(ts int64) ([]*schema.MetricData) {
  metrics := make([]*schema.MetricData, keyCount)

  for i := 1; i <= keyCount; i++ {
    metrics[i-1] = &schema.MetricData{
      Name:     fmt.Sprintf(keyPrefix + "%d", i),
      Metric:   fmt.Sprintf(keyPrefix + "%d", i),
      OrgId:    i,
      Value:    0,
      Unit:     "ms",
      Mtype:    "gauge",
      Tags:     []string{"some_tag", "ok", "k:2"},
      Time:     ts,
    }
  }
  return metrics
}
