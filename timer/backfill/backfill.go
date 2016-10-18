package backfill

import (
	"flag"
	"fmt"
	"time"

	mod "github.com/OOM-Killer/fakemetrics_ng/timer/module"
	dur "github.com/raintank/dur"
	gc "github.com/rakyll/globalconf"
)

var Module *mod.ModuleT = &mod.ModuleT{
	"backfill",
	func(id int) mod.Timer {
		return &Backfill{
			id,
			0,
			make(chan time.Time),
		}
	},
	RegisterFlagSet,
}

type Backfill struct {
	agentId      int
	internalTime int64
	ticker       chan time.Time
}

type LongDurationFlag uint32

var (
	interval    int
	startOffset LongDurationFlag
)

func (ld *LongDurationFlag) Set(value string) error {
	*ld = LongDurationFlag(dur.MustParseUsec("long duration", value))
	return nil
}

func (ld *LongDurationFlag) String() string {
	return fmt.Sprintf("%d", ld)
}

func RegisterFlagSet() {
	flags := flag.NewFlagSet("backfill", flag.ExitOnError)
	flags.IntVar(&interval, "interval", 1000, "the metric interval")
	flags.Var(&startOffset, "start-offset", "offset to start backfill from")
	gc.Register("backfill", flags)
}

func (bf *Backfill) GetInterval() int {
	return interval
}

func (bf *Backfill) loop() {
	for {
		t := time.Unix(bf.internalTime, 0)
		bf.ticker <- t
	}
}

func (bf *Backfill) GetTicker() <-chan time.Time {
	go bf.loop()
	return bf.ticker
}

func (bf *Backfill) GetTimestamp() int64 {
	if bf.internalTime == 0 {
		bf.internalTime = time.Now().Unix() - int64(startOffset)
	} else {
		bf.internalTime = bf.internalTime + int64(interval/1000)
	}
	return bf.internalTime
}
