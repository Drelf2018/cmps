package cmps

import "reflect"

type Value struct {
	reflect.Value
}

func (v *Value) Any(i int) any {
	return v.Field(i).Interface()
}

func NewValue(x any) *Value {
	return &Value{reflect.ValueOf(x)}
}
