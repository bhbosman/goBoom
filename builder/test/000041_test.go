package test

import (
	"reflect"
	"testing"
)

type MyType struct {
}

func TestReflect(t *testing.T) {
	//v := reflect.ValueOf(fmt.Sprintf)
	//r := v.Call([]reflect.Value{reflect.ValueOf("dddd")})
	//println(r[0].String())

	field01 := reflect.StructField{
		Name:      "Field01",
		PkgPath:   "",
		Type:      reflect.TypeOf("string"),
		Tag:       "",
		Offset:    0,
		Index:     nil,
		Anonymous: false,
	}
	field02 := reflect.StructField{
		Name:      "Field02",
		PkgPath:   "",
		Type:      reflect.TypeOf(0),
		Tag:       "",
		Offset:    0,
		Index:     nil,
		Anonymous: false,
	}
	fields := []reflect.StructField{field01, field02}
	structType := reflect.StructOf(fields)
	instance := reflect.New(structType)
	println(instance.Type(), instance.Elem().String())
	instance.Elem().Field(0).SetString("2222")
	instance.Elem().Field(1).SetInt(2222)
	println(instance.Elem().String())

}
