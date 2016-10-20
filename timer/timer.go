package timer

import (
	"fmt"
	"time"
)

type Timer interface {
	GetInterval() int
	GetTicker() <-chan time.Time
	GetTimestamp() int64
}
type tConstructor func(int)(Timer)

var modules map[string]tConstructor = make(map[string]tConstructor)
var regFlags []func()

func RegFlags() {
	for _, reg := range regFlags {
		reg()
	}
}

func Get(name string, id int) Timer {
	mod, ok := modules[name]
	if !ok {
		panic(fmt.Sprintf("failed to find timer %s", name))
	}
	return mod(id)
}
