package functions

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestTranslateToInt(t *testing.T) {
	m := make(map[FuncKey]reflect.Value)
	RegisterTranslateToInt(m)
	t.Run("Translate int8 To int", func(t *testing.T) {
		key := FuncKey{
			Name: "int",
			DataType: reflect.FuncOf(
				[]reflect.Type{reflect.TypeOf(int8(0))},
				[]reflect.Type{reflect.TypeOf(0)},
				false,
			),
		}
		r, ok := m[key]
		assert.True(t, ok)
		v := r.Call([]reflect.Value{reflect.ValueOf(int8(0))})
		assert.Same(t, v[0].Type(), reflect.TypeOf(0))
	})
	t.Run("Translate int16 To int", func(t *testing.T) {
		key := FuncKey{
			Name: "int",
			DataType: reflect.FuncOf(
				[]reflect.Type{reflect.TypeOf(int16(0))},
				[]reflect.Type{reflect.TypeOf(0)},
				false,
			),
		}
		r, ok := m[key]
		assert.True(t, ok)
		v := r.Call([]reflect.Value{reflect.ValueOf(int16(0))})
		assert.Same(t, v[0].Type(), reflect.TypeOf(0))
	})
	t.Run("Translate int32 To int", func(t *testing.T) {
		key := FuncKey{
			Name: "int",
			DataType: reflect.FuncOf(
				[]reflect.Type{reflect.TypeOf(int32(0))},
				[]reflect.Type{reflect.TypeOf(0)},
				false,
			),
		}
		r, ok := m[key]
		assert.True(t, ok)
		v := r.Call([]reflect.Value{reflect.ValueOf(int32(0))})
		assert.Same(t, v[0].Type(), reflect.TypeOf(0))
	})
	t.Run("Translate int64 To int", func(t *testing.T) {
		key := FuncKey{
			Name: "int",
			DataType: reflect.FuncOf(
				[]reflect.Type{reflect.TypeOf(int64(0))},
				[]reflect.Type{reflect.TypeOf(0)},
				false,
			),
		}
		r, ok := m[key]
		assert.True(t, ok)
		v := r.Call([]reflect.Value{reflect.ValueOf(int64(0))})
		assert.Same(t, v[0].Type(), reflect.TypeOf(0))
	})
	t.Run("Translate int To int", func(t *testing.T) {
		key := FuncKey{
			Name: "int",
			DataType: reflect.FuncOf(
				[]reflect.Type{reflect.TypeOf(0)},
				[]reflect.Type{reflect.TypeOf(0)},
				false,
			),
		}
		r, ok := m[key]
		assert.True(t, ok)
		v := r.Call([]reflect.Value{reflect.ValueOf(0)})
		assert.Same(t, v[0].Type(), reflect.TypeOf(0))
	})
}
