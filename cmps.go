package cmps

import (
	"cmp"
	"slices"
)

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
	switch t := any(x).(type) {
	default:
		return packages.Get(x).Compare(x, y, nil)
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
	}
}

func Search[S ~[]T, T any](x S, target T) (int, bool) {
	return slices.BinarySearchFunc(x, target, Compare)
}

func Insert[S ~[]T, T any](x S, target T) S {
	i, _ := Search(x, target)
	return slices.Insert(x, i, target)
}

func Slice[S ~[]E, E any](x S) {
	slices.SortFunc(x, Compare)
}

func SliceWithGroup[S ~[]E, E any](x S, g *Group) {
	slices.SortFunc(x, func(a, b E) int { return packages.Get(a).Compare(a, b, g) })
}

type Map[K comparable, V any] struct {
	K K `cmps:"114;groups:key"`
	V V `cmps:"514;groups:value"`
}

func makeMaps[M ~map[K]V, K comparable, V any](m M) []Map[K, V] {
	r := make([]Map[K, V], len(m))
	i := 0
	for k, v := range m {
		r[i] = Map[K, V]{k, v}
		i++
	}
	return r
}

func Keys[M ~map[K]V, K comparable, V any](m M) []Map[K, V] {
	r := makeMaps(m)
	SliceWithGroup(r, &Group{Reserve: []string{"key"}})
	return r
}

func Values[M ~map[K]V, K comparable, V any](m M) []Map[K, V] {
	r := makeMaps(m)
	SliceWithGroup(r, &Group{Reserve: []string{"value"}})
	return r
}

func Contains[M ~map[K]V, K comparable, V any](m M, f func(K, V) bool) bool {
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}
