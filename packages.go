package cmps

import (
	"reflect"
)

type Packages map[string]map[string]Type

func (p Packages) Set(pkg, name string, typ reflect.Type, fields ...Field) {
	if p.Contain(pkg, name) {
		return
	}
	c := Type{Type: typ, Fields: fields, pkg: pkg, name: name}
	c.Parse()
	p[pkg][name] = c
}

func (p Packages) Contain(pkg, name string) bool {
	types := p[pkg]
	if types == nil {
		types = make(map[string]Type)
		p[pkg] = types
		return false
	}
	_, ok := types[name]
	return ok
}

func (p Packages) Get(v any) Type {
	typ := check(v)
	pkg, name := typ.PkgPath(), typ.Name()
	p.Set(pkg, name, typ)
	return p[pkg][name]
}
