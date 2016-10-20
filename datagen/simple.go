package datagen

import (
	"flag"
	"fmt"

	"gopkg.in/raintank/schema.v1"

	gc "github.com/rakyll/globalconf"
)

type Simple struct {
	id int
}

var (
	sKeyCount  int
	sKeyPrefix string
)

func init() {
	modules["simple"] = sNew
	regFlags = append(regFlags, sRegFlags)
}

func sNew(id int) (Datagen) {
	return &Simple{id}
}

func sRegFlags() {
	flags := flag.NewFlagSet("simple", flag.ExitOnError)
	flags.IntVar(&sKeyCount, "key-count", 100, "number of keys to generate")
	flags.StringVar(&sKeyPrefix, "key-prefix", "some.key.", "prefix for keys")
	gc.Register("simple", flags)
}

func (s *Simple) GetData(ts int64) []*schema.MetricData {
	metrics := make([]*schema.MetricData, sKeyCount)

	for i := 0; i < sKeyCount; i++ {
		metrics[i] = &schema.MetricData{
			Name:   fmt.Sprintf(sKeyPrefix+"%d.%d", s.id, i),
			Metric: fmt.Sprintf(sKeyPrefix+"%d.%d", s.id, i),
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
