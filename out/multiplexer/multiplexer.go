package multiplexer

import (
	"gopkg.in/raintank/schema.v1"

	mod "github.com/OOM-Killer/fakemetrics_ng/out/module"
)

type Multiplexer struct {
	in       chan *schema.MetricData
	outs     []mod.OutIface
	outChans []chan *schema.MetricData
}

func (m *Multiplexer) Start() {
	for _, out := range m.outs {
		out.Start()
	}
}

func (m *Multiplexer) Put(metric *schema.MetricData) {
	for _, out := range m.outs {
		out.Put(metric)
	}
}

func (m *Multiplexer) AddOut(out mod.OutIface) {
	m.outs = append(m.outs, out)
}

func (m *Multiplexer) RegisterFlagSet() {}
