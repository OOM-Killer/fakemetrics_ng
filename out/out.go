package out

import (
	"fmt"

	carbon "github.com/OOM-Killer/fakemetrics_ng/out/carbon"
	mod "github.com/OOM-Killer/fakemetrics_ng/out/module"
	mp "github.com/OOM-Killer/fakemetrics_ng/out/multiplexer"
)

var (
	moduleMap []*mod.ModuleT = []*mod.ModuleT{
		carbon.Module,
	}
)

func RegisterFlagSets() {
	for _, o := range moduleMap {
		o.RegFlags()
	}
}

func GetInstance(seek string) mod.OutIface {
	for _, o := range moduleMap {
		if o.Name == seek {
			return o.Init()
		}
	}
	panic(fmt.Sprintf("failed to find output %s", seek))
}

func GetMultiInstance(names []string) mod.OutIface {
	m := mp.Multiplexer{}
	for _, name := range names {
		m.AddOut(GetInstance(name))
	}
	return &m
}
