package timer

import (
	"flag"
	"fmt"
	"time"

	dur "github.com/raintank/dur"
	gc "github.com/rakyll/globalconf"
)

type Backfill struct {
	id           int
	internalTime int64
	ticker       chan time.Time
}

type LongDurationFlag uint32

var (
	bfInterval  int
	startOffset LongDurationFlag
)

func init() {
	modules["backfill"] = bfNew
	regFlags = append(regFlags, bfRegFlags)
}

func bfNew(id int, agents int) Timer {
	return &Backfill{id, 0, make(chan time.Time)}
}

func bfRegFlags() {
	flags := flag.NewFlagSet("backfill", flag.ExitOnError)
	flags.IntVar(&bfInterval, "interval", 1000, "the metric interval")
	flags.Var(&startOffset, "start-offset", "offset to start backfill from")
	gc.Register("backfill", flags)
}

func (ld *LongDurationFlag) Set(value string) error {
	*ld = LongDurationFlag(dur.MustParseUsec("long duration", value))
	return nil
}

func (ld *LongDurationFlag) String() string {
	return fmt.Sprintf("%d", ld)
}

func (bf *Backfill) GetInterval() int {
	return bfInterval
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
		bf.internalTime = bf.internalTime + int64(bfInterval/1000)
	}
	return bf.internalTime
}
