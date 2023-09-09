package cmps

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var packages = new(Packages)

type Packages struct {
	sync.Map
}

func (p *Packages) Load(key any) (value *Type, ok bool) {
	v, ok := p.Map.Load(key)
	if ok {
		value = v.(*Type)
	}
	return
}

func (p *Packages) Set(pkg, name string, typ reflect.Type, names ...string) *Type {
	key := pkg + "|" + name
	if p, ok := p.Load(key); ok {
		return p
	}
	c := Type{Type: typ}
	if len(names) != 0 {
		c.Scan(names)
		c.toIndex(pkg, name)
	} else {
		c.findIndex(pkg, name)
	}
	p.Store(key, &c)
	return &c
}

func (p *Packages) Get(v any) *Type {
	typ := checkValue(v)
	return p.Set(typ.PkgPath(), typ.Name(), typ)
}

func Show() {
	current_pkg := ""
	packages.Range(func(key, value any) bool {
		keys := strings.Split(key.(string), "|")
		pkg := keys[0]
		name := keys[1]
		typ := value.(*Type)
		if current_pkg != pkg {
			current_pkg = pkg
			println(pkg)
		}
		println("├", name)
		m := len(typ.Fields)
		for _, f := range typ.Fields {
			if m == 1 {
				print("│ └ ", f.Name)
			} else {
				print("│ ├ ", f.Name)
			}
			if len(f.Options.Groups) != 0 {
				fmt.Printf(" %v\n", f.Options.Groups)
			} else {
				println()
			}
			m--
		}
		return true
	})

}
