package main

import (
	"flag"
	"strings"

	gc "github.com/rakyll/globalconf"

	"github.com/OOM-Killer/fakemetrics_ng/agents"
	"github.com/OOM-Killer/fakemetrics_ng/datagen"
	"github.com/OOM-Killer/fakemetrics_ng/out"
	"github.com/OOM-Killer/fakemetrics_ng/timer"
)

type stringListFlags []string

func (f *stringListFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func (f *stringListFlags) String() string {
	return strings.Join(*f, ", ")
}

var (
	timerMod   string
	dataGenMod string
	outMod     stringListFlags
)

func setupConfig() {
	timer.RegFlags()
	datagen.RegFlags()
	out.RegisterFlagSets()
	agents.RegisterFlagSets()

	conf, err := gc.NewWithOptions(
		&gc.Options{
			Filename: *confFile,
		})
	if err != nil {
    panic(err)
	}

	moduleFlags := flag.NewFlagSet("modules", flag.ExitOnError)
	moduleFlags.StringVar(
		&timerMod,
		"timer",
		"realtime",
		"the name of the timer module")
	moduleFlags.StringVar(
		&dataGenMod,
		"data-gen",
		"simple",
		"the name of the data generator module")
	moduleFlags.Var(
		&outMod,
		"output",
		"name of the output module, can be specified multiple times")
	gc.Register("modules", moduleFlags)
	conf.ParseAll()
}
