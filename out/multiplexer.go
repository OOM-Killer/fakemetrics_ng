package out

import (
	"gopkg.in/raintank/schema.v1"
)

type Multiplexer struct {
	Id       int
	in       chan *schema.MetricData
	outs     []Out
	outChans []chan *schema.MetricData
}

func init() {
	modules["multiplexer"] = mpNew
}

func mpNew(id int) (Out) {
	m := &Multiplexer{}
	m.Id = id
	return m
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

func (m *Multiplexer) AddOut(out Out) {
	m.outs = append(m.outs, out)
}
