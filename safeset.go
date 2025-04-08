package goset

import "sync"

type safeSet[T comparable] struct {
	ncs *unsafeSet[T]
	sync.RWMutex
}

// Interface guard to assert concrete type:safeSet adheres to Set interface.
var _ Set[string] = (*safeSet[string])(nil)

func newSafeSet[T comparable](size int) *safeSet[T] {
	return &safeSet[T]{
		ncs: newUnsafeSet[T](size),
	}
}

func (s *safeSet[T]) Add(v T) {
	s.Lock()
	defer s.Unlock()
	s.ncs.Add(v)
}

func (s *safeSet[T]) Remove(v T) {
	s.Lock()
	defer s.Unlock()
	s.ncs.Remove(v)
}

func (s *safeSet[T]) Contains(v T) bool {
	s.RLock()
	defer s.RUnlock()

	return s.ncs.Contains(v)
}

func (s *safeSet[T]) Size() int {
	s.RLock()
	defer s.RUnlock()

	return s.ncs.Size()
}

func (s *safeSet[T]) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()

	return s.ncs.IsEmpty()
}

func (s *safeSet[T]) IsEqual(other Set[T]) bool {
	o, _ := other.(*safeSet[T])

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	return s.ncs.IsEqual(o.ncs)
}

func (s *safeSet[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.ncs.Clear()
}

func (s *safeSet[T]) ToSlice() []T {
	s.RLock()
	defer s.RUnlock()

	return s.ncs.ToSlice()
}

func (s *safeSet[T]) Clone() Set[T] {
	s.RLock()
	defer s.RUnlock()
	clone, _ := s.ncs.Clone().(*unsafeSet[T])

	return &safeSet[T]{
		ncs: clone,
	}
}

func (s *safeSet[T]) Union(other Set[T]) Set[T] {
	o, _ := other.(*safeSet[T])

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	union, _ := s.ncs.Union(o.ncs).(*unsafeSet[T])

	return &safeSet[T]{
		ncs: union,
	}
}

func (s *safeSet[T]) Intersection(other Set[T]) Set[T] {
	o, _ := other.(*safeSet[T])

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	intersection, _ := s.ncs.Intersection(o.ncs).(*unsafeSet[T])

	return &safeSet[T]{
		ncs: intersection,
	}
}

func (s *safeSet[T]) Difference(other Set[T]) Set[T] {
	o, _ := other.(*safeSet[T])

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	difference, _ := s.ncs.Difference(o.ncs).(*unsafeSet[T])

	return &safeSet[T]{
		ncs: difference,
	}
}
