package cmps

import (
	"encoding/json"
	"sync"

	"golang.org/x/exp/slices"
)

type Slice[T any] struct {
	// I stands for Items
	I  []T
	rw sync.RWMutex
}

func (s *Slice[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.I)
}

func (s *Slice[T]) Index(t T) int {
	s.rw.RLock()
	defer s.rw.RUnlock()
	if i, ok := Search(s.I, t); ok {
		return i
	}
	return -1
}

func (s *Slice[T]) Search(t T) (zero T) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	if i, ok := Search(s.I, t); ok {
		return s.I[i]
	}
	return
}

func (s *Slice[T]) Insert(t T) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.I = Insert(s.I, t)
}

func (s *Slice[T]) Delete(t T) bool {
	s.rw.Lock()
	defer s.rw.Unlock()
	if i, ok := Search(s.I, t); ok {
		s.I = slices.Delete(s.I, i, i+1)
		return true
	}
	return false
}

func (s *Slice[T]) Sort() {
	s.rw.Lock()
	defer s.rw.Unlock()
	Sort(s.I)
}

func NewSlice[T any](items ...T) *Slice[T] {
	return &Slice[T]{I: items}
}
