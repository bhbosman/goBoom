package functions

import "reflect"

func translateIntToInt(m map[FuncKey]reflect.Value) {
	key := FuncKey{
		Name: "int",
		DataType: reflect.FuncOf(
			[]reflect.Type{reflect.TypeOf(int(0))},
			[]reflect.Type{reflect.TypeOf(int(0))},
			false,
		),
	}
	m[key] = reflect.MakeFunc(
		key.DataType,
		func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(int(args[0].Int()))}
		},
	)
}

func translateInt8ToInt(m map[FuncKey]reflect.Value) {
	key := FuncKey{
		Name: "int",
		DataType: reflect.FuncOf(
			[]reflect.Type{reflect.TypeOf(int8(0))},
			[]reflect.Type{reflect.TypeOf(int(0))},
			false,
		),
	}
	m[key] = reflect.MakeFunc(
		key.DataType,
		func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(int(args[0].Int()))}
		},
	)
}

func translateInt16ToInt(m map[FuncKey]reflect.Value) {
	key := FuncKey{
		Name: "int",
		DataType: reflect.FuncOf(
			[]reflect.Type{reflect.TypeOf(int16(0))},
			[]reflect.Type{reflect.TypeOf(int(0))},
			false,
		),
	}
	m[key] = reflect.MakeFunc(
		key.DataType,
		func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(int(args[0].Int()))}
		},
	)
}
func translateInt32ToInt(m map[FuncKey]reflect.Value) {
	key := FuncKey{
		Name: "int",
		DataType: reflect.FuncOf(
			[]reflect.Type{reflect.TypeOf(int32(0))},
			[]reflect.Type{reflect.TypeOf(int(0))},
			false,
		),
	}
	m[key] = reflect.MakeFunc(
		key.DataType,
		func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(int(args[0].Int()))}
		},
	)
}
func translateInt64ToInt(m map[FuncKey]reflect.Value) {
	key := FuncKey{
		Name: "int",
		DataType: reflect.FuncOf(
			[]reflect.Type{reflect.TypeOf(int64(0))},
			[]reflect.Type{reflect.TypeOf(int(0))},
			false,
		),
	}
	m[key] = reflect.MakeFunc(
		key.DataType,
		func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(int(args[0].Int()))}
		},
	)
}

func RegisterTranslateToInt(m map[FuncKey]reflect.Value) {
	translateInt8ToInt(m)
	translateInt16ToInt(m)
	translateInt32ToInt(m)
	translateInt64ToInt(m)
	translateIntToInt(m)
}
