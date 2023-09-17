package cmps

import (
	"encoding/json"
	"slices"
	"sync"
)

type SafeSlice[T any] struct {
	// I stands for Items
	I  []T
	rw sync.RWMutex
}

func (s *SafeSlice[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.I)
}

func (s *SafeSlice[T]) Index(t T) int {
	s.rw.RLock()
	defer s.rw.RUnlock()
	if i, ok := Search(s.I, t); ok {
		return i
	}
	return -1
}

func (s *SafeSlice[T]) Search(t T) (zero T) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	if i, ok := Search(s.I, t); ok {
		return s.I[i]
	}
	return
}

func (s *SafeSlice[T]) Insert(t T) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.I = Insert(s.I, t)
}

func (s *SafeSlice[T]) Delete(t T) {
	s.rw.Lock()
	defer s.rw.Unlock()
	if i, ok := Search(s.I, t); ok {
		s.I = slices.Delete(s.I, i, i+1)
	}
}

func (s *SafeSlice[T]) Sort() {
	s.rw.Lock()
	defer s.rw.Unlock()
	Slice(s.I)
}
