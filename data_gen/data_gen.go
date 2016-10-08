package data_gen

import (
  "fmt"

  "gopkg.in/raintank/schema.v1"
)

type DataGen interface {
  RegisterFlagSet()
  GetData() (*schema.MetricData)
  GetName() (string)
}

type DataGenFactory struct {
  dataGens []DataGen
}

func NewFactory() DataGenFactory {
  inst := DataGenFactory{}
  inst.initDataGens()
  inst.registerFlagSets()
  return inst
}

func (f *DataGenFactory) initDataGens() {
  f.dataGens = []DataGen{
    //&simple.Simple{}
  }
}

func (f *DataGenFactory) registerFlagSets() {
  for _, dataGen := range f.dataGens {
    dataGen.RegisterFlagSet()
  }
}

func (f *DataGenFactory) GetDataGen(seek string) (DataGen) {
  f.initDataGens()
  for _,dataGen := range f.dataGens {
    if (dataGen.GetName() == seek) {
      return dataGen
    }
  }
  panic(fmt.Sprintf("could not find data generator %s", seek))
}
