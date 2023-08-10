package cmps

import (
	"cmp"
	"errors"
	"reflect"
)

var (
	packages  = make(Packages)
	NotStruct = errors.New("v's element type must be struct")
	SameType  = errors.New("the type of y must same as x")
)

func check(v any) reflect.Type {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		panic(NotStruct)
	}
	return typ
}

func Compare(x, y any, options ...Type) int {
	switch x := x.(type) {
	case bool:
		if x == y.(bool) {
			return 0
		}
		if x {
			return +1
		}
		return -1
	case int:
		return cmp.Compare(x, y.(int))
	case int8:
		return cmp.Compare(x, y.(int8))
	case int16:
		return cmp.Compare(x, y.(int16))
	case int32:
		return cmp.Compare(x, y.(int32))
	case int64:
		return cmp.Compare(x, y.(int64))
	case uint:
		return cmp.Compare(x, y.(uint))
	case uint8:
		return cmp.Compare(x, y.(uint8))
	case uint16:
		return cmp.Compare(x, y.(uint16))
	case uint32:
		return cmp.Compare(x, y.(uint32))
	case uint64:
		return cmp.Compare(x, y.(uint64))
	case uintptr:
		return cmp.Compare(x, y.(uintptr))
	case float32:
		return cmp.Compare(x, y.(float32))
	case float64:
		return cmp.Compare(x, y.(float64))
	case string:
		return cmp.Compare(x, y.(string))
	default:
		var t Type
		if len(options) == 0 {
			t = packages.Get(x)
			if !t.Equal(y) {
				panic(SameType)
			}
		} else {
			t = options[0]
		}
		xv, yv := NewValue(x), NewValue(y)
		for _, f := range t.Fields {
			result := Compare(xv.Any(f.Index), yv.Any(f.Index), t.Child(f.Name))
			if result != 0 {
				return result
			}
		}
		return 0
	}
}
