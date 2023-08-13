package cmps

import (
	"fmt"
	"reflect"
)

type Ordered interface {
	OrderBy() []string
}

var packages = make(Packages)

type Packages map[string]map[string]Type

func (p Packages) Set(pkg, name string, typ reflect.Type, names ...string) *Type {
	if p.Contain(pkg, name) {
		c := p[pkg][name]
		return &c
	}
	c := Type{Type: typ}
	if len(names) != 0 {
		c.Scan(names)
		c.toIndex(pkg, name)
	} else {
		c.findIndex(pkg, name)
	}
	p[pkg][name] = c
	return &c
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

func (p Packages) Get(v any) *Type {
	typ := checkValue(v)
	if v, ok := v.(Ordered); ok {
		return p.Set(typ.PkgPath(), typ.Name(), typ, v.OrderBy()...)
	}
	return p.Set(typ.PkgPath(), typ.Name(), typ)
}

func Show() {
	for pkg, types := range packages {
		println(pkg)
		l := len(types)
		for name, typ := range types {
			if l == 1 {
				println("└", name)
			} else {
				println("├", name)
			}
			chn := "│"
			if l == 1 {
				chn = " "
			}
			m := len(typ.Fields)
			for _, f := range typ.Fields {
				if m == 1 {
					print(chn+" └ ", f.Name)
				} else {
					print(chn+" ├ ", f.Name)
				}
				if len(f.Options.Groups) != 0 {
					fmt.Printf(" %v\n", f.Options.Groups)
				} else {
					println()
				}
				m--
			}
			l--
		}
	}
}
