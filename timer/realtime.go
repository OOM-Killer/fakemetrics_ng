package timer

import (
	"flag"
	"time"

	gc "github.com/rakyll/globalconf"
)


type Realtime struct {
	id int
}

var (
	rtInterval int
)

func init() {
	modules["realtime"] = rtNew
	regFlags = append(regFlags, rtRegFlags)
}

func rtNew(id int) (Timer) {
	return &Realtime{id}
}

func rtRegFlags() {
	flags := flag.NewFlagSet("realtime", flag.ExitOnError)
	flags.IntVar(&rtInterval, "interval", 100, "the metric interval")
	gc.Register("realtime", flags)
}

func (r *Realtime) GetTicker() <-chan time.Time {
	return time.NewTicker(time.Duration(rtInterval) * time.Millisecond).C
}

func (r *Realtime) GetTimestamp() int64 {
	return time.Now().Unix()
}

func (r *Realtime) GetInterval() int {
	return rtInterval
}
