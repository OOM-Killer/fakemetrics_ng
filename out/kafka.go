package out

import (
	"flag"
  "fmt"
  "time"
  "encoding/binary"

  "github.com/Shopify/sarama"
	"gopkg.in/raintank/schema.v1"

	gc "github.com/rakyll/globalconf"
)

var (
  kZkHost    string
  kZkPort    int
  kMsgCount  int
  kFlushTime int
  kTopic     string
  kCodec     string
)

type Kafka struct {
  Id       int
  producer sarama.SyncProducer
  in       chan *schema.MetricData
}

func init() {
  modules["kafka"] = kNew
  regFlags = append(regFlags, kRegFlags)
}

func kNew(id int) Out {
  k := Kafka{}
  k.Id = id
  return &k
}

func kRegFlags() {
  flags := flag.NewFlagSet("kafka", flag.ExitOnError)
  flags.StringVar(&kZkHost, "zookeeper-host", "localhost", "zookeeper hostname")
  flags.IntVar(&kZkPort, "zookeeper-port", 2181, "zookeeper port")
  flags.StringVar(&kCodec, "codec", "none", "compression: none|gzip|snappy")
  flags.IntVar(&kMsgCount, "msg-metric-count", 100, "max metrics per message")
  flags.IntVar(&kFlushTime, "msg-flush-time", 100, "max time a metric gets queue before flushing")
  flags.StringVar(&kTopic, "kafka-topic", "mdm", "the kafka topic to send metrics to")
  gc.Register("kafka", flags)
}

func (k *Kafka) Put(metric *schema.MetricData) {
	if k.in == nil {
		panic("can't accept data before starting")
	}

  k.in <- metric
}

func (k *Kafka) Start() {
  k.in = make(chan *schema.MetricData)

  config := sarama.NewConfig()
  config.Producer.RequiredAcks = sarama.WaitForAll
  config.Producer.Retry.Max = 10
  config.Producer.Compression = getCompression(kCodec)
  producer, err := sarama.NewSyncProducer(
    []string{fmt.Sprintf("%s:%d", kZkHost, kZkPort)},
    config,
  )

  if err != nil {
    panic(err)
  }

  k.producer = producer

  go k.Loop()
}

func (k *Kafka) Loop() error {
	var t = time.NewTicker(time.Duration(kFlushTime) * time.Millisecond)
  var data []byte
  pos := 0
  payload := make([]*sarama.ProducerMessage, kMsgCount)

  for {
    select {
    case <-t.C:
      if (pos > 0) {
        k.flush(payload[:pos])
        pos = 0
      }
    case metric := <-k.in:
      data, err := metric.MarshalMsg(data[:])
      if err != nil {
        return err
      }

      key := make([]byte, 8)
      binary.LittleEndian.PutUint32(key, uint32(metric.OrgId))

      payload[pos] = &sarama.ProducerMessage{
        Key:   sarama.ByteEncoder(key),
        Topic: kTopic,
        Value: sarama.ByteEncoder(data),
      }

      if (pos >= kMsgCount) {
        k.flush(payload)
        pos = 0
      }
    }
  }
}

func (k *Kafka) flush(payload []*sarama.ProducerMessage) {
  k.producer.SendMessages(payload)
}

func getCompression(codec string) sarama.CompressionCodec {
	switch codec {
	case "none":
		return sarama.CompressionNone
	case "gzip":
		return sarama.CompressionGZIP
	case "snappy":
		return sarama.CompressionSnappy
	default:
		panic(fmt.Sprintf("unknown compression codec %q", codec))
	}
}
