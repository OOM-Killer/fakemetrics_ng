package buffered_writer

import (
  "bytes"
  "fmt"
  "time"

  "gopkg.in/raintank/schema.v1"
)

type funFlush func(*[]byte)

type BufferedWriter struct {
  FlushInterval int
  MetricsPerFlush int
  FlushCB funFlush
  in chan *schema.MetricData
  buffer []*schema.MetricData
  bufferPos int
}

func (b *BufferedWriter) GetChan() (chan *schema.MetricData) {
  if b.in == nil {
    b.in = make(chan *schema.MetricData, b.MetricsPerFlush)
  }
  return b.in
}

func (b *BufferedWriter) Loop() {
  var t = time.NewTicker(time.Duration(b.FlushInterval) * time.Millisecond)
  b.bufferPos = 0
  b.buffer = make([]*schema.MetricData, b.MetricsPerFlush)

  for {
    select {
    case <-t.C:
      b.flush()
    case metric := <-b.in:
      if (b.bufferPos >= b.MetricsPerFlush) {
        <-t.C
        b.flush()
      }
      b.buffer[b.bufferPos] = metric
      b.bufferPos++
    }
  }
}

func (b *BufferedWriter) flush() {
  var m *schema.MetricData
  fmt.Println(fmt.Sprintf("flushing buffer of length %d", b.bufferPos))
  buf := bytes.NewBufferString("")

  for i := 0; i < b.bufferPos; i++ {
    m = b.buffer[i]
    buf.WriteString(fmt.Sprintf("%s %f %d\n", m.Name, m.Value, m.Time))
  }
  b.bufferPos = 0

  bytes := buf.Bytes()
  b.FlushCB(&bytes)
}
