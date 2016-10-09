package factory

import (
  "fmt"
)

type Module interface {
  RegisterFlagSet()
  GetName() (string)
}

type Factory struct {
  instances []Module
}

func (f *Factory) RegisterModule(mod Module) {
  f.instances = append(f.instances, mod)
}

func (f *Factory) RegisterFlagSets() {
  for _,inst := range f.instances {
    inst.RegisterFlagSet()
  }
}

func (f *Factory) GetInstance(seek string)(Module) {
  for _,inst := range f.instances {
    if (inst.GetName() == seek) {
      return inst
    }
  }
  panic(fmt.Sprintf("could not find module %s", seek))
}
