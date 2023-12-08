package cmps

import (
	"golang.org/x/exp/slices"

	"golang.org/x/exp/constraints"
)

// Compare returns
//
//	-1 if x is less than y,
//	 0 if x equals y,
//	+1 if x is greater than y.
//
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
func StdCompare[T constraints.Ordered](x, y T) int {
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
func isNaN[T constraints.Ordered](x T) bool {
	return x != x
}

func compare[T constraints.Ordered](x, y any, t T) int {
	return StdCompare(x.(T), y.(T))
}

func CmpBool(x, y any) int {
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
	return CompareUnsafe(x, y)
}

func CompareUnsafe(x, y any) int {
	switch t := any(x).(type) {
	default:
		return Tags(ref.Get(x)).Compare(x, y)
	case bool:
		return CmpBool(x, y)
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
	return slices.BinarySearchFunc(x, target, Compare[T])
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

func Sort[S ~[]E, E any](x S) {
	slices.SortFunc(x, Compare[E])
}
