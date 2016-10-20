package datagen

import (
	"fmt"

	"gopkg.in/raintank/schema.v1"
)

type Datagen interface {
	GetData(int64) []*schema.MetricData
}
type dgConstructor func(int) Datagen

var modules map[string]dgConstructor = make(map[string]dgConstructor)
var regFlags []func()

func RegFlags() {
	for _, reg := range regFlags {
		reg()
	}
}

func Get(name string, id int) Datagen {
	mod, ok := modules[name]
	if !ok {
		panic(fmt.Sprintf("could not find datagen %q", name))
	}
	return mod(id)
}
