package cmps

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func init() {
	packages.Get(new(Field))
}

func checkValue(v any) reflect.Type {
	if r, ok := checkType(reflect.TypeOf(v)); ok {
		return r
	}
	panic(NotStruct)
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

type values struct {
	Value1 reflect.Value
	Value2 reflect.Value
}

func (v *values) Any(i int) (any, any) {
	return v.Value1.Field(i).Interface(), v.Value2.Field(i).Interface()
}

func newValues(x, y any) *values {
	v1 := reflect.ValueOf(x)
	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}
	v2 := reflect.ValueOf(y)
	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}
	return &values{v1, v2}
}

type Field struct {
	Index    int
	Name     string
	Cmps     float64 `cmps:"114"`
	isStruct bool
}

type Type struct {
	reflect.Type
	Fields []*Field
	pkg    string
	name   string
}

func (t Type) Compare(x, y any) int {
	vs := newValues(x, y)
	for _, f := range t.Fields {
		result := 0
		if f.isStruct {
			typ := packages[t.pkg][t.name+"."+f.Name]
			result = typ.Compare(vs.Any(f.Index))
		} else {
			result = Compare(vs.Any(f.Index))
		}
		if result != 0 {
			return result
		}
	}
	return 0
}

func (t *Type) Equal(v any) bool {
	typ := checkValue(v)
	return t.pkg == typ.PkgPath() && t.name == typ.Name()
}

func (t *Type) findIndex() {
	length := t.Type.NumField()
	for i := 0; i < length; i++ {
		v := t.Type.Field(i)
		tag := v.Tag.Get("cmps")
		if tag == "" {
			continue
		}
		tags := strings.Split(tag, ";")
		cmps, err := strconv.ParseFloat(tags[0], 64)
		if err != nil {
			panic(fmt.Errorf("The tag of %v is not a number(%v)", v.Name, err))
		}
		_, ok := checkType(v.Type)
		if ok {
			if len(tags) == 1 {
				t.NewType(v)
			} else {
				names := strings.Split(tags[1], ",")
				fields := make([]*Field, len(names))
				for i, name := range names {
					fields[i] = &Field{Name: name, Cmps: float64(i)}
				}
				t.NewType(v, fields...)
			}
		}
		t.Fields = append(t.Fields, &Field{Index: i, Name: v.Name, Cmps: cmps, isStruct: ok})
	}
	// 自举
	Slice(t.Fields)
}

func (t *Type) toIndex() {
	for _, field := range t.Fields {
		v, ok := t.FieldByName(field.Name)
		if !ok {
			panic(fmt.Errorf("field: %v was not found", field.Name))
		}
		field.Index = v.Index[0]
		if _, ok := checkType(v.Type); ok {
			field.isStruct = true
			t.NewType(v)
		}
	}
}

func (t *Type) parse() {
	if len(t.Fields) != 0 {
		t.toIndex()
	} else {
		t.findIndex()
	}
}

func (t *Type) NewType(v reflect.StructField, fields ...*Field) {
	c := Type{Type: v.Type, Fields: fields, pkg: t.pkg, name: t.name + "." + v.Name}
	c.parse()
	packages[c.pkg][c.name] = c
}
