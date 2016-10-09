package simple

import (
  "flag"
  "fmt"

  "gopkg.in/raintank/schema.v1"

  gc "github.com/rakyll/globalconf"
)

type Simple struct {}

var (
  keyCount int
  keyPrefix string
)

func (s *Simple) RegisterFlagSet() {
  flags := flag.NewFlagSet(s.GetName(), flag.ExitOnError)
  flags.IntVar(&keyCount, "key_count", 100, "number of keys to generate")
  flags.StringVar(&keyPrefix, "key_prefix", "some.key.", "prefix for keys")
  gc.Register(s.GetName(), flags)
}

func (s *Simple) GetName() (string) {
  return "simple"
}

func (s *Simple) GetData(ts int64) (*schema.MetricData) {
  return &schema.MetricData{
    Name:     fmt.Sprintf(keyPrefix + "%d", 1),
    Metric:   "a.c.a.f",
    OrgId:    1,
    Value:    0,
    Unit:     "ms",
    Mtype:    "gauge",
    Tags:     []string{"some_tag", "ok", "k:2"},
    Time:     ts,
  }
}
