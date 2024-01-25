package cmps

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Drelf2018/TypeGo/Reflect"
)

func ParseTag(tag string) (m map[string]string) {
	m = make(map[string]string)
	if len(tag) == 0 {
		return
	}
	for _, item := range strings.Split(tag, ";") {
		items := strings.SplitN(item, ":", 2)
		if len(items) > 1 {
			m[strings.ToLower(items[0])] = items[1]
		} else {
			m["cmps"] = items[0]
		}
	}
	return
}

type Tag struct {
	Reflect.Field
	Cmps   float64 `cmps:"1"`
	Order  int
	Fields Tags
}

func (t Tag) String() string {
	buf := bytes.NewBufferString(", [")
	l := len(t.Fields)
	for i, field := range t.Fields {
		buf.WriteString(field.String())
		if i != l-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return fmt.Sprintf("Tag(%v, %v, cmps=%v, order=%v%v)", t.Name, t.Index, t.Cmps, t.Order, buf.String())
}

func (Tag) Insert(field Reflect.Field, ok bool, fields []*Tag, array *[]*Tag) {
	tag := &Tag{
		Field:  field,
		Cmps:   -1,
		Order:  1,
		Fields: fields,
	}

	for key, value := range ParseTag(field.Tag) {
		switch key {
		case "cmps":
			tag.Cmps, _ = strconv.ParseFloat(value, 64)
		case "fields":
			m := make(map[string]Tag)
			for _, v := range fields {
				m[v.Name] = *v
			}
			names := strings.Split(value, ",")
			fields = make([]*Tag, 0, len(names))
			for i, name := range names {
				t := m[name]
				t.Cmps = float64(i)
				fields = append(fields, &t)
			}
			tag.Fields = fields
		case "order":
			if strings.ToLower(value) == "desc" {
				tag.Order = -1
			}
		}
	}

	*array = Insert(*array, tag)
}

func (t *Tag) Compare(x, y any) int {
	if len(t.Fields) == 0 {
		return Compare(x, y)
	}
	return t.Fields.Compare(x, y)
}

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

type Tags []*Tag

func (t Tags) Compare(x, y any) int {
	v := equip(x, y)
	for _, f := range t {
		if f.Cmps < 0 {
			continue
		}
		if r := f.Compare(v.any(f.Index)); r != 0 {
			return f.Order * r
		}
	}
	return 0
}

var ref = Reflect.NewMap[Reflect.TagParser[Tag], []*Tag]("cmps")
