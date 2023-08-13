package cmps

import "slices"

type Group struct {
	Reserve []string
	Omit    []string
}

func (g Group) single(group string) bool {
	if slices.Contains(g.Omit, group) {
		return false
	}
	if len(g.Reserve) == 0 {
		return true
	}
	return slices.Contains(g.Reserve, group)
}

func (g Group) multiple(groups []string) bool {
	for _, group := range groups {
		if !g.single(group) {
			return false
		}
	}
	return true
}
