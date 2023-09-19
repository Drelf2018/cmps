package cmps

import (
	"golang.org/x/exp/slices"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Compare returns
//
//	-1 if x is less than y,
//	 0 if x equals y,
//	+1 if x is greater than y.
//
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
func cmpCompare[T Ordered](x, y T) int {
	xNaN := isNaN(x)
	yNaN := isNaN(y)
	if xNaN && yNaN {
		return 0
	}
	if xNaN || x < y {
		return -1
	}
	if yNaN || x > y {
		return +1
	}
	return 0
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T Ordered](x T) bool {
	return x != x
}

func compare[T Ordered](x, y any, t T) int {
	return cmpCompare(x.(T), y.(T))
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
	return slices.BinarySearchFunc(x, target, func(t1, t2 T) int { return Compare(t1, t2) })
}

func SearchFunc[F any, S ~[]T, T *F](x S, f func(T)) (int, bool) {
	var target T = new(F)
	f(target)
	return Search(x, target)
}

func Insert[S ~[]T, T any](x S, target T) S {
	i, _ := Search(x, target)
	return slices.Insert(x, i, target)
}

func Delete[S ~[]T, T any](x S, target T) S {
	if i, ok := Search(x, target); ok {
		return slices.Delete(x, i, i+1)
	}
	return x
}

func Slice[S ~[]E, E any](x S) {
	slices.SortFunc(x, func(a, b E) int { return Compare(a, b) })
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
