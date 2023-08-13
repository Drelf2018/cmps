package cmps

import (
	"errors"
	"reflect"
)

func checkValue(v any) reflect.Type {
	if r, ok := checkType(reflect.TypeOf(v)); ok {
		return r
	}
	panic(errors.New("v's element type must be struct"))
}

func checkType(vt reflect.Type) (reflect.Type, bool) {
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	if vt.Kind() == reflect.Struct {
		return vt, true
	}
	return nil, false
}
