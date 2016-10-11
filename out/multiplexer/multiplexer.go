package multiplexer


import (
  //"fmt"

  "gopkg.in/raintank/schema.v1"

  iface "github.com/OOM-Killer/fakemetrics_ng/out/iface"
)

type Multiplexer struct {
  in chan *schema.MetricData
  outs []iface.OutIface
  outChans []chan *schema.MetricData
}

func (m *Multiplexer) Start() {
  for _,out := range m.outs {
    out.Start()
  }
}

func (m *Multiplexer) Put(metric *schema.MetricData) {
  for _,out := range m.outs {
    out.Put(metric)
  }
}

func (m *Multiplexer) AddOut(out iface.OutIface) {
  m.outs = append(m.outs, out)
}

func (m *Multiplexer) GetName() (string) {
  return "multiplexer"
}

func (m *Multiplexer) RegisterFlagSet() {}
