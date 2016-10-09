package carbon

import (
  "fmt"
  "flag"
  "time"

  "gopkg.in/raintank/schema.v1"

  gc "github.com/rakyll/globalconf"
)

type Carbon struct {
  in chan *schema.MetricData
  buffer []*schema.MetricData
  bufferPos int
}

var (
  flushInterval int
  metricsPerFlush int
)

func (c *Carbon) RegisterFlagSet() {
  flags := flag.NewFlagSet(c.GetName(), flag.ExitOnError)
  flags.IntVar(&flushInterval, "flush-interval", 100, "the metric interval")
  flags.IntVar(&metricsPerFlush, "metrics-per-flush", 10, "the metric interval")
  gc.Register(c.GetName(), flags)
}

func (c *Carbon) GetName() (string) {
  return "carbon"
}

func (c *Carbon) GetChan() (chan *schema.MetricData) {
  if (c.in == nil) {
    panic ("can't provide channel before starting")
  }
  return c.in
}

func (c *Carbon) Start() {
  c.in = make(chan *schema.MetricData, metricsPerFlush)
  go c.loop()
}

func (c *Carbon) loop() {
  var t = time.NewTicker(time.Duration(flushInterval) * time.Millisecond)
  c.bufferPos = 0
  c.buffer = make([]*schema.MetricData, metricsPerFlush)

  for {
    select {
    case <-t.C:
      c.flush()
    case metric := <-c.in:
      if (c.bufferPos >= metricsPerFlush) {
        <-t.C
        c.flush()
      }
      c.buffer[c.bufferPos] = metric
      c.bufferPos++
    }
  }
}

func (c *Carbon) flush() {
  fmt.Println(fmt.Sprintf("buffer length %d", c.bufferPos))
  // do some flushing
  c.bufferPos = 0
}
