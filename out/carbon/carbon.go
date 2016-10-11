package carbon

import (
  "fmt"
  "flag"
  "log"
  "time"
  "net"

  "gopkg.in/raintank/schema.v1"

  bw "github.com/OOM-Killer/fakemetrics_ng/out/buffered_writer"
  gc "github.com/rakyll/globalconf"
)

type Carbon struct {
  in chan *schema.MetricData
  bw *bw.BufferedWriter
  conn net.Conn
}

var (
  flushInterval int
  metricsPerFlush int
  writeBufferSize int
  host string
  port int
)

func (c *Carbon) RegisterFlagSet() {
  flags := flag.NewFlagSet(c.GetName(), flag.ExitOnError)
  flags.IntVar(&flushInterval, "flush-interval", 100, "the metric interval")
  flags.IntVar(&metricsPerFlush, "metrics-per-flush", 10, "the metric interval")
  flags.IntVar(&writeBufferSize, "write-buffer-size", 1000, "write buffer size")
  flags.StringVar(&host, "host", "localhost", "carbon host name")
  flags.IntVar(&port, "port", 2003, "carbon port")
  gc.Register(c.GetName(), flags)
}

func (c *Carbon) GetName() (string) {
  return "carbon"
}

func (c *Carbon) Put(metric *schema.MetricData) {
  if (c.in == nil) {
    panic ("can't accept data before starting")
  }

  select {
  case c.in <- metric:
  default:
    log.Println("write buffer full. output is slow or buffer too small")
  }
}

func (c *Carbon) Start() {
  c.bw = &bw.BufferedWriter{}
  c.bw.FlushInterval = flushInterval
  c.bw.MetricsPerFlush = metricsPerFlush
  c.bw.WriteBufferSize = writeBufferSize
  c.bw.FlushCB = c.flush
  c.in = c.bw.GetChan()
  c.connect()
  go c.bw.Loop()
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

func (c *Carbon) flush(buf *[]byte) {
  _, err := c.conn.Write(*buf)
  if err != nil {
    // desperate attempt to prevent losing the data in the buffer
    c.connect()
    c.conn.Write(*buf)
  }
}
