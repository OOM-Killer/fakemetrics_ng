package out

import (
	"fmt"
	"gopkg.in/raintank/schema.v1"
)

type Out interface {
	Start()
	Put(*schema.MetricData)
}
type oConstructor func(int) Out

var modules map[string]oConstructor = make(map[string]oConstructor)
var regFlags []func()

func RegFlags() {
	for _, reg := range regFlags {
		reg()
	}
}

func Get(name string, id int) Out {
	mod, ok := modules[name]
	if !ok {
		panic(fmt.Sprintf("failed to find output %s", name))
	}
	return mod(id)
}

func GetMulti(names []string, id int) Out {
	m := Multiplexer{}
	for _, name := range names {
		m.AddOut(Get(name, id))
	}
	return &m
}
