package bean

import (
	"github.com/facebookgo/inject"
	"reflect"
)

var (
	injectGraph *inject.Graph = &inject.Graph{}
)

func Provide(objects ...*inject.Object) error {
	return injectGraph.Provide(objects...)
}

func Populate() error {
	return injectGraph.Populate()
}

func GetObjectByName(name string) interface{} {
	for _, o := range injectGraph.Objects() {
		if o.Name == name {
			return o.Value
		}

	}
	return nil
}

func GetObjectByType(value interface{}) (res []interface{}) {
	rv := reflect.TypeOf(value)
	for _, o := range injectGraph.Objects() {
		if reflect.TypeOf(o.Value) == rv {
			res = append(res, o.Value)
		}
	}
	return
}

func GetOneObjectByType(value interface{}) interface{} {
	rv := reflect.TypeOf(value)
	for _, o := range injectGraph.Objects() {
		if reflect.TypeOf(o.Value) == rv {
			return o.Value
		}
	}
	return nil
}
