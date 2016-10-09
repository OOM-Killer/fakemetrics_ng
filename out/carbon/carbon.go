package carbon

import (
  "fmt"
  "flag"
  "time"
  "net"
  "bytes"

  "gopkg.in/raintank/schema.v1"

  gc "github.com/rakyll/globalconf"
)

type Carbon struct {
  in chan *schema.MetricData
  buffer []*schema.MetricData
  bufferPos int
  conn net.Conn
}

var (
  flushInterval int
  metricsPerFlush int
  host string
  port int
)

func (c *Carbon) RegisterFlagSet() {
  flags := flag.NewFlagSet(c.GetName(), flag.ExitOnError)
  flags.IntVar(&flushInterval, "flush-interval", 100, "the metric interval")
  flags.IntVar(&metricsPerFlush, "metrics-per-flush", 10, "the metric interval")
  flags.StringVar(&host, "host", "localhost", "carbon host name")
  flags.IntVar(&port, "port", 2003, "carbon port")
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
  c.connect()
  go c.loop()
}

func (c *Carbon) connect() {
  for {
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
    if err == nil {
      c.conn = conn
      break
    } else {
      fmt.Println("failed to connect to carbon, retrying")
      time.Sleep(100 * time.Millisecond)
    }
  }
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
  var m *schema.MetricData
  fmt.Println(fmt.Sprintf("flushing buffer of length %d", c.bufferPos))
  buf := bytes.NewBufferString("")

  for i := 0; i < c.bufferPos; i++ {
    m = c.buffer[i]
    buf.WriteString(fmt.Sprintf("%s %f %d\n", m.Name, m.Value, m.Time))
  }
  c.bufferPos = 0

  _, err := c.conn.Write(buf.Bytes())
  if err != nil {
    // desperate attempt to prevent losing the data in the buffer
    c.connect()
    c.conn.Write(buf.Bytes())
  }
}
