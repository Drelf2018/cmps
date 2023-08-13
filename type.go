package cmps

import (
	"fmt"
	"reflect"
)

type Type struct {
	reflect.Type
	Fields
}

func (t *Type) Compare(x, y any, g *Group) int {
	if t == nil {
		return Compare(x, y)
	}
	v := equip(x, y)
	for _, f := range t.Fields {
		if g != nil && !g.multiple(f.Groups) {
			continue
		}
		v0, v1 := v.any(f.Index)
		switch f.Type.Compare(v0, v1, g) {
		case 1:
			return 1
		case -1:
			return -1
		}
	}
	return 0
}

func (t *Type) findIndex(pkg, name string) {
	length := t.Type.NumField()
	for i := 0; i < length; i++ {
		v := t.Type.Field(i)
		tag := v.Tag.Get("cmps")
		if tag == "" {
			continue
		}
		field := Field{Index: i, Name: v.Name}
		field.Options.parse(tag)
		field.Type = t.ComplexNew(pkg, name, v, field.Options.Fields...)
		t.Fields = append(t.Fields, &field)
	}
	// 自举
	if len(t.Fields) > 1 {
		Slice(t.Fields)
	}
}

func (t *Type) toIndex(pkg, name string) {
	for _, field := range t.Fields {
		v, ok := t.FieldByName(field.Name)
		if !ok {
			panic(fmt.Errorf("field: %v was not found", field.Name))
		}
		field.Index = v.Index[0]
		field.Type = t.ComplexNew(pkg, name, v)
	}
}

func (t *Type) EasyNew(pkg, name string, v reflect.StructField, fields ...string) *Type {
	c := Type{Type: v.Type}
	if len(fields) != 0 {
		c.Scan(fields)
	}
	vname := name + "." + v.Name
	if len(c.Fields) != 0 {
		c.toIndex(pkg, vname)
	} else {
		c.findIndex(pkg, vname)
	}
	packages[pkg][vname] = c
	return &c
}

func (t *Type) ComplexNew(pkg, name string, v reflect.StructField, fields ...string) *Type {
	if vt, ok := checkType(v.Type); ok {
		if vt.PkgPath() == "" {
			return t.EasyNew(pkg, name, v, fields...)
		}
		nv := reflect.New(vt).Interface()
		if nv, ok := nv.(Ordered); ok {
			return packages.Set(vt.PkgPath(), vt.Name(), vt, nv.OrderBy()...)
		} else {
			return packages.Set(vt.PkgPath(), vt.Name(), vt, fields...)
		}
	}
	return nil
}
