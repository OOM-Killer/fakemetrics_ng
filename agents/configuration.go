package agents

import (
  "flag"

  gc "github.com/rakyll/globalconf"
)

var (
  agentCount int
  offsets string
)

func RegisterFlagSets() {
  flags := flag.NewFlagSet("multiagent", flag.ExitOnError)
  flags.IntVar(
    &agentCount,
    "agent-count",
    1000,
    "number of agents to run")
  flags.StringVar(
    &offsets,
    "offsets",
    "random",
    "how to distribute the agent offsets")
  gc.Register("multiagent", flags)
}
