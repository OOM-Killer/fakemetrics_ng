package timer

import (
	"fmt"
	"flag"
	"time"
	"math/rand"

	gc "github.com/rakyll/globalconf"
)

type Flexible struct {
	id       int
	agents   int
	interval int
}

var (
	fMaxInterval        int
	fMinInterval        int
	randDist bool
)

func init() {
	modules["flexible"] = fNew
	regFlags = append(regFlags, fRegFlags)
	rand.Seed(time.Now().Unix())
}

func fNew(id int, agents int) Timer {
	var interval int
	if randDist {
		// a random number between min and max
		interval = rand.Intn(fMaxInterval - fMinInterval) + fMinInterval
	} else {
		// intervals are evenly distributed between min and max
		interval = ((fMaxInterval - fMinInterval) / agents * id) + fMinInterval
	}
	fmt.Println(fmt.Sprintf("starting agent %d with interval %d", id, interval))
	return &Flexible{id, agents, interval}
}

func fRegFlags() {
	flags := flag.NewFlagSet("flexible", flag.ExitOnError)
	flags.IntVar(&fMaxInterval, "max-interval", 3000, "maximum metric interval")
	flags.IntVar(&fMinInterval, "min-interval", 100, "minimum metric interval")
	flags.BoolVar(&randDist, "random-distribution", false, "randomize distribution of intervals")
	gc.Register("flexible", flags)
}

func (f *Flexible) GetTicker() <-chan time.Time {
	return time.NewTicker(time.Duration(f.interval) * time.Millisecond).C
}

func (f *Flexible) GetTimestamp() int64 {
	return time.Now().Unix()
}

func (f *Flexible) GetInterval() int {
	return f.interval
}
