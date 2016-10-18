package key_changer

import (
	"flag"
	"fmt"

	"gopkg.in/raintank/schema.v1"

	mod "github.com/OOM-Killer/fakemetrics_ng/data_gen/module"
	gc "github.com/rakyll/globalconf"
)

var (
	pointsPerKey int
	keyCount     int
	keyPrefix    string
	syncSwitch   bool
)

var Module *mod.ModuleT = &mod.ModuleT{
	"key-changer",
	func(id int) mod.DataGen { return New(id) },
	RegisterFlagSet,
}

type KeyChanger struct {
	agentId   int
	keyPoints []int
	currKey   []int
}

func New(id int) *KeyChanger {
	initValue := 0
	keyPoints := make([]int, keyCount)
	currKey := make([]int, keyCount)

	for i := 0; i < keyCount; i++ {
		currKey[i] = 0
		keyPoints[i] = initValue
		if !syncSwitch {
			initValue++
		}
	}

	return &KeyChanger{id, keyPoints, currKey}
}

func RegisterFlagSet() {
	flags := flag.NewFlagSet("key-changer", flag.ExitOnError)
	flags.IntVar(&pointsPerKey, "points-per-key", 10, "number of points per key")
	flags.IntVar(&keyCount, "key-count", 100, "number of keys to generate")
	flags.StringVar(&keyPrefix, "key-prefix", "some.key", "prefix for keys")
	flags.BoolVar(&syncSwitch, "sync-switch", true, "change all keys at once")
	gc.Register("key-changer", flags)
}

func (kc *KeyChanger) GetData(ts int64) []*schema.MetricData {
	metrics := make([]*schema.MetricData, keyCount)

	for i := 0; i < keyCount; i++ {
		name := fmt.Sprintf(keyPrefix+"%d.%d.%d", kc.agentId, i, kc.currKey[i])
		metrics[i] = &schema.MetricData{
			Name:   name,
			Metric: name,
			OrgId:  i,
			Value:  0,
			Unit:   "ms",
			Mtype:  "gauge",
			Tags:   []string{"some_tag", "ok", "k:2"},
			Time:   ts,
		}

		kc.keyPoints[i]++

		if kc.keyPoints[i]%pointsPerKey == 0 {
			kc.keyPoints[i] = 0
			kc.currKey[i]++
		}
	}

	return metrics
}
