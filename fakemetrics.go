package main

import (
	"flag"

	"github.com/OOM-Killer/fakemetrics_ng/agents"
	"github.com/OOM-Killer/fakemetrics_ng/data_gen"
	"github.com/OOM-Killer/fakemetrics_ng/out"
	"github.com/OOM-Killer/fakemetrics_ng/timer"
)

var (
	confFile = flag.String("config",
		"fakemetrics.ini",
		"configuration file path")
)

func main() {
	flag.Parse()
	timer.RegisterFlagSets()
	data_gen.RegisterFlagSets()
	out.RegisterFlagSets()
	agents.RegisterFlagSets()

	setupConfig()

	a := agents.New(timerMod, dataGenMod, outMod)
	a.Run()
}
