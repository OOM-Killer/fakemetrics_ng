package datagen

import (
	"flag"
	"fmt"

	"gopkg.in/raintank/schema.v1"

	gc "github.com/rakyll/globalconf"
)

var (
	sKeyCount  int
	sKeyPrefix string
)

type Simple struct {
	agentid int
}

func init() {
	modules["simple"] = NewSimple
	regFlags = append(regFlags, RegFlagsSimple)
}

func NewSimple(agentid int) (Datagen) {
	return &Simple{agentid}
}

func RegFlagsSimple() {
	flags := flag.NewFlagSet("simple", flag.ExitOnError)
	flags.IntVar(&sKeyCount, "key-count", 100, "number of keys to generate")
	flags.StringVar(&sKeyPrefix, "key-prefix", "some.key.", "prefix for keys")
	gc.Register("simple", flags)
}

func (s *Simple) GetData(ts int64) []*schema.MetricData {
	metrics := make([]*schema.MetricData, sKeyCount)

	for i := 0; i < sKeyCount; i++ {
		metrics[i] = &schema.MetricData{
			Name:   fmt.Sprintf(sKeyPrefix+"%d.%d", s.agentid, i),
			Metric: fmt.Sprintf(sKeyPrefix+"%d.%d", s.agentid, i),
			OrgId:  i,
			Value:  0,
			Unit:   "ms",
			Mtype:  "gauge",
			Tags:   []string{"some_tag", "ok", "k:2"},
			Time:   ts,
		}
	}
	return metrics
}
