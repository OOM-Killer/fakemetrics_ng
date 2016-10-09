package multiplexer


import (
  "gopkg.in/raintank/schema.v1"

  iface "github.com/OOM-Killer/fakemetrics_ng/out/iface"
)

type Multiplexer struct {
  in chan *schema.MetricData
  outs []iface.OutIface
}

func (m *Multiplexer) Start() {
  var outChans []chan *schema.MetricData
  for _,out := range m.outs {
    outChans = append(outChans, out.GetChan())
  }

  m.in = make(chan *schema.MetricData)
  for {
    metric := <-m.in
    for _,c := range outChans {
      c<-metric
    }
  }
}

func (m *Multiplexer) AddOut(out *iface.OutIface) {
  m.outs = append(m.outs, *out)
}

func (m *Multiplexer) GetChan() (chan *schema.MetricData) {
  if (m.in == nil) {
    panic ("can't provide channel before starting")
  }
  return m.in
}

func (m *Multiplexer) GetName() (string) {
  return "multiplexer"
}

func (m *Multiplexer) RegisterFlagSet() {}
