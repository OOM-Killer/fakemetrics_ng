package main

import (
  "fmt"
  "flag"
  //"os"
  //"time"
  "strings"

  "github.com/rakyll/globalconf"
  "github.com/OOM-Killer/fakemetrics_ng/timer"
)


type outModFlags []string

func (f *outModFlags) Set(value string) error {
  *f = append(*f, value)
  return nil
}

func (f *outModFlags) String() string {
  return strings.Join(*f, ", ")
}

var (
  // flags
  confFile    = flag.String("config",
                            "fakemetrics.ini",
                            "configuration file path")
  timerMod    = flag.String("timer",
                            "realtime",
                            "the name of the timer module")
  dataGenMod  = flag.String("data-gen",
                            "default",
                            "the name of the data generator module")
  outMod outModFlags

  // Module factories
  //timerFactory *timer.TimerFactory
)

func main() {
  flag.Var(&outMod, "output",
    "name of the output module, can be specified multiple times")
  flag.Parse()

  conf, err := globalconf.NewWithOptions(
    &globalconf.Options{
      Filename: *confFile,
    })
  if err != nil {
    panic("error with configuration file")
  }

  timerFactory := timer.NewFactory()

  fmt.Println("parsing now")
  conf.ParseAll()

  timer := timerFactory.GetTimer(*timerMod)
  timer.PrintInterval()

  /*fmt.Println(*timerMod)
  fmt.Println(*dataGenMod)
  fmt.Println(strings.Join(outMod, ","))*/


  //interval := time.Duration(1000) * time.Millisecond
  tick := timer.GetTicker()

  for range tick.C {
    fmt.Println("tick")
  }

}
