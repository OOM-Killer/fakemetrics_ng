package agents

import (
  "flag"

  gc "github.com/rakyll/globalconf"
)

var (
  agentCount int
  offsets string
  slowIncrease bool
  launchInterval int
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
  flags.BoolVar(
    &slowIncrease,
    "slow-increase",
    true,
    "increase number of agents slowly")
  flags.IntVar(
    &launchInterval,
    "launch-interval",
    100,
    "interval between launching agents in ms")
  gc.Register("multiagent", flags)
}
