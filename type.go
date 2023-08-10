package cmps

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type Field struct {
	Index int
	Name  string
}

type Type struct {
	reflect.Type
	Fields []Field
	pkg    string
	name   string
}

func (t *Type) Child(name string) Type {
	return packages[t.pkg][t.name+"."+name]
}

func (t *Type) Need(name string) int {
	if len(t.Fields) == 0 {
		return -2
	}
	return slices.IndexFunc(t.Fields, func(f Field) bool { return f.Name == name })
}

func (t *Type) Equal(v any) bool {
	typ := check(v)
	return t.pkg == typ.PkgPath() && t.name == typ.Name()
}

func (t *Type) ForField(f func(i int, f float64, v reflect.StructField, field ...string)) {
	length := t.Type.NumField()
	for i := 0; i < length; i++ {
		v := t.Type.Field(i)
		pos := t.Need(v.Name)
		tag := v.Tag.Get("cmps")
		// 如果父类没有指定顺序并且自己也没写 tag 就跳过
		if pos == -1 && tag == "" {
			continue
		}
		tags := strings.Split(tag, ";")
		var arg float64
		if pos >= 0 {
			arg = float64(pos)
		} else {
			var err error
			arg, err = strconv.ParseFloat(tags[0], 64)
			if err != nil {
				panic(fmt.Errorf("The tag of %v is not a number(%v)", v.Name, err))
			}
		}
		if len(tags) == 1 {
			f(i, arg, v)
		} else {
			field := strings.Split(tags[1], ",")
			f(i, arg, v, field...)
		}
	}
}

func (t *Type) Parse() {
	idxMap := make(map[float64]Field)
	t.ForField(func(i int, f float64, v reflect.StructField, field ...string) {
		idx, ok := idxMap[f]
		if ok {
			panic(fmt.Errorf("cmps:\"%v\" has been used by field %v", f, idx))
		}
		if v.Type.Kind() == reflect.Struct {
			var fields = make([]Field, v.Type.NumField())
			for i, f := range field {
				fields[i] = Field{Name: f}
			}
			packages.Set(t.pkg, t.name+"."+v.Name, v.Type, fields...)
		}
		idxMap[f] = Field{Index: i, Name: v.Name}
	})
	t.Fields = SortMap(idxMap)
}
