package agents

import (
	"time"

	"github.com/OOM-Killer/fakemetrics_ng/datagen"
	out "github.com/OOM-Killer/fakemetrics_ng/out/module"
	"github.com/OOM-Killer/fakemetrics_ng/timer"
)

type Agent struct {
	timer   timer.Timer
	datagen datagen.Datagen
	out     out.OutIface
	offset  int
}

func (a *Agent) Run() {
	time.Sleep(time.Duration(a.offset))

	a.out.Start()
	tick := a.timer.GetTicker()
	for range tick {
		go a.doTick()
	}
}

func (a *Agent) doTick() {
	metrics := a.datagen.GetData(a.timer.GetTimestamp())
	for _, m := range metrics {
		a.out.Put(m)
	}
}
