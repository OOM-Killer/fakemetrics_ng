package buffered_writer

import (
  "time"

  "gopkg.in/raintank/schema.v1"
)

type funFlush func([]*schema.MetricData)

type BufferedWriter struct {
  FlushInterval int
  MetricsPerFlush int
  FlushCB funFlush
  WriteBufferSize int
  in chan *schema.MetricData
  buffer []*schema.MetricData
  bufferPos int
}

func (b *BufferedWriter) GetChan() (chan *schema.MetricData) {
  if b.in == nil {
    b.in = make(chan *schema.MetricData, b.WriteBufferSize)
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
  b.FlushCB(b.buffer[:b.bufferPos])
  b.bufferPos = 0
}
