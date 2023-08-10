package cmps

import (
	"cmp"
	"slices"
)

// 傻逼 go1.21 把 exp 的 maps 包添加到了标准库但是又不把这个函数加进来 是不是脑子让驴踢了
//
// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// 根据键排序返回值列表
func SortMap[M ~map[K]V, K cmp.Ordered, V any](m M) []V {
	v := make([]V, len(m))
	keys := Keys(m)
	slices.Sort(keys)
	for i, k := range keys {
		v[i] = m[k]
	}
	return v
}
