package agents

import (
  "time"
  "fmt"
  "math/rand"

  "github.com/OOM-Killer/fakemetrics_ng/timer"
  "github.com/OOM-Killer/fakemetrics_ng/data_gen"
  "github.com/OOM-Killer/fakemetrics_ng/out"
)

type Agents struct {
  time string
  dataGen string
  out []string
  agents []*Agent
}

func New(t string, dg string, o []string) (*Agents) {
  return &Agents{t, dg, o, make([]*Agent, agentCount)}
}

func (a *Agents) Run() {
  var os int

  for i := 0; i < agentCount; i++ {
    timer := timer.GetInstance(a.time)
    switch offsets {
    case "none":
      os = 0
    case "even":
      os = 1e9 * timer.GetInterval() / agentCount
    case "random":
      os = 1e9 * timer.GetInterval() / (rand.Intn(1e9) + 1)
    default:
      panic(fmt.Sprintf("invalid offset mode %s", offsets))
    }

    a.agents[i] = &Agent{
      timer,
      data_gen.GetInstance(a.dataGen),
      out.GetMultiInstance(a.out),
      os,
    }
  }

  // separate the two loops to not add initialization time to offsets
  var t *time.Ticker

  for i := 0; i < agentCount; i++ {
    go a.agents[i].Run()

    if (slowIncrease) {
      if (t == nil) {
        t = time.NewTicker(time.Duration(launchInterval) * time.Millisecond)
      }
      fmt.Println(fmt.Sprintf("agent count is %d", i + 1))
      <-t.C
    }
  }

  <-make(chan bool)
}
