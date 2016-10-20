package out

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"gopkg.in/raintank/schema.v1"

	gc "github.com/rakyll/globalconf"
)

var (
	cFlushInterval   int
	cMetricsPerFlush int
	cWriteBufferSize int
	cBlockOnWrite    bool
	cHost            string
	cPort            int
)

type Carbon struct {
	Id   int
	in   chan *schema.MetricData
	bw   *BufferedWriter
	conn net.Conn
}

func init() {
	modules["carbon"] = cNew
	regFlags = append(regFlags, cRegFlags)
}

func cNew(id int) (Out) {
	c := &Carbon{}
	c.Id = id
	return c
}

func cRegFlags() {
	flags := flag.NewFlagSet("carbon", flag.ExitOnError)
	flags.IntVar(&cFlushInterval, "flush-interval", 100, "the metric interval")
	flags.IntVar(&cMetricsPerFlush, "metrics-per-flush", 10, "the metric interval")
	flags.IntVar(&cWriteBufferSize, "write-buffer-size", 1000, "write buffer size")
	flags.StringVar(&cHost, "host", "localhost", "carbon host name")
	flags.IntVar(&cPort, "port", 2003, "carbon port")
	flags.BoolVar(&cBlockOnWrite, "block-on-write", false, "block on slow write")
	gc.Register("carbon", flags)
}

func (c *Carbon) Put(metric *schema.MetricData) {
	if c.in == nil {
		panic("can't accept data before starting")
	}

	if cBlockOnWrite {
		c.in <- metric
	} else {
		select {
		case c.in <- metric:
		default:
			log.Println("write buffer full. output is slow or buffer too small")
		}
	}
}

func (c *Carbon) Start() {
	c.bw = &BufferedWriter{}
	c.bw.FlushInterval = cFlushInterval
	c.bw.MetricsPerFlush = cMetricsPerFlush
	c.bw.WriteBufferSize = cWriteBufferSize
	c.bw.FlushCB = c.flush
	c.in = c.bw.GetChan()
	c.connect()
	go c.bw.Loop()
}

func (c *Carbon) connect() {
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", cHost, cPort))
		if err == nil {
			c.conn = conn
			break
		} else {
			fmt.Println("failed to connect to carbon, retrying")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (c *Carbon) flush(metrics []*schema.MetricData) {
	buf := bytes.NewBufferString("")

	for _, m := range metrics {
		buf.WriteString(fmt.Sprintf("%s %f %d\n", m.Name, m.Value, m.Time))
	}

	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		// desperate attempt to prevent losing the data in the buffer
		c.connect()
		c.conn.Write(buf.Bytes())
	}
}
