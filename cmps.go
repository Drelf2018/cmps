package cmps

import (
	"cmp"
	"errors"
	"reflect"
	"slices"
)

var (
	packages  = make(Packages)
	NotStruct = errors.New("v's element type must be struct")
	SameType  = errors.New("the type of y must same as x")
)

type Packages map[string]map[string]Type

func (p Packages) Set(pkg, name string, typ reflect.Type) Type {
	if p.Contain(pkg, name) {
		return p[pkg][name]
	}
	c := Type{Type: typ, pkg: pkg, name: name}
	c.parse()
	p[pkg][name] = c
	return c
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
	typ := checkValue(v)
	return p.Set(typ.PkgPath(), typ.Name(), typ)
}

func compare[T cmp.Ordered](x, y any, t T) int {
	return cmp.Compare(x.(T), y.(T))
}

func cmpbool(x, y any) int {
	a := x.(bool)
	if a == y.(bool) {
		return 0
	}
	if a {
		return +1
	}
	return -1
}

func Compare[T any](x, y T) int {
	var t T = x
	switch t := any(t).(type) {
	case bool:
		return cmpbool(x, y)
	case int:
		return compare(x, y, t)
	case int8:
		return compare(x, y, t)
	case int16:
		return compare(x, y, t)
	case int32:
		return compare(x, y, t)
	case int64:
		return compare(x, y, t)
	case uint:
		return compare(x, y, t)
	case uint8:
		return compare(x, y, t)
	case uint16:
		return compare(x, y, t)
	case uint32:
		return compare(x, y, t)
	case uint64:
		return compare(x, y, t)
	case uintptr:
		return compare(x, y, t)
	case float32:
		return compare(x, y, t)
	case float64:
		return compare(x, y, t)
	case string:
		return compare(x, y, t)
	default:
		typ := packages.Get(x)
		if !typ.Equal(y) {
			panic(SameType)
		}
		return typ.Compare(x, y)
	}
}

func Slice[S ~[]E, E any](x S) {
	slices.SortFunc(x, Compare)
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	Slice(r)
	return r
}

func KeysToValues[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, len(m))
	for i, k := range Keys(m) {
		r[i] = m[k]
	}
	return r
}

func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	Slice(r)
	return r
}

type reversedMap[K comparable, V any] struct {
	K K
	V V `cmps:"514"`
}

func ValuesToKeys[M ~map[K]V, K comparable, V any](m M) []K {
	rm := make([]reversedMap[K, V], 0, len(m))
	for k, v := range m {
		rm = append(rm, reversedMap[K, V]{k, v})
	}
	Slice(rm)
	s := make([]K, len(m))
	for i, r := range rm {
		s[i] = r.K
	}
	return s
}
