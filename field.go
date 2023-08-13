package cmps

import (
	"reflect"
	"strings"
)

type values [2]reflect.Value

func (v *values) any(i int) (any, any) {
	return v[0].Field(i).Interface(), v[1].Field(i).Interface()
}

func equip(x, y any) values {
	v0 := reflect.ValueOf(x)
	if v0.Kind() == reflect.Ptr {
		v0 = v0.Elem()
	}
	v1 := reflect.ValueOf(y)
	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}
	return values{v0, v1}
}

type Field struct {
	Index int
	Name  string
	Options
	*Type
}

func (f Field) OrderBy() []string {
	return []string{"Options"}
}

type Fields []*Field

func (fs *Fields) Read(s string) {
	fs.Scan(strings.Split(s, ","))
}

func (fs *Fields) Scan(names []string) {
	*fs = make([]*Field, len(names))
	for i, name := range names {
		(*fs)[i] = &Field{Name: name, Options: Options{Cmps: float64(i)}}
	}
}
