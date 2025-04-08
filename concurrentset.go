package goset

import "sync"

type concurrentSet[T comparable] struct {
	ncs *noconcurrentSet[T]
	sync.RWMutex
}

// Interface guard to assert concrete type:concurrentSet adheres to Set interface.
var _ Set[string] = (*concurrentSet[string])(nil)

func newConcurrentSet[T comparable](size int) *concurrentSet[T] {
	return &concurrentSet[T]{
		ncs: newNoConcurrentSet[T](size),
	}
}

func (s *concurrentSet[T]) Add(v T) {
	s.Lock()
	defer s.Unlock()
	s.ncs.Add(v)
}

func (s *concurrentSet[T]) Remove(v T) {
	s.Lock()
	defer s.Unlock()
	s.ncs.Remove(v)
}

func (s *concurrentSet[T]) Contains(v T) bool {
	s.RLock()
	defer s.RUnlock()

	return s.ncs.Contains(v)
}

func (s *concurrentSet[T]) Size() int {
	s.RLock()
	defer s.RUnlock()

	return s.ncs.Size()
}

func (s *concurrentSet[T]) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()

	return s.ncs.IsEmpty()
}

func (s *concurrentSet[T]) IsEqual(other Set[T]) bool {
	o, _ := other.(*concurrentSet[T])

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	return s.ncs.IsEqual(o.ncs)
}

func (s *concurrentSet[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.ncs.Clear()
}

func (s *concurrentSet[T]) ToSlice() []T {
	s.RLock()
	defer s.RUnlock()

	return s.ncs.ToSlice()
}

func (s *concurrentSet[T]) Clone() Set[T] {
	s.RLock()
	defer s.RUnlock()
	clone, _ := s.ncs.Clone().(*noconcurrentSet[T])

	return &concurrentSet[T]{
		ncs: clone,
	}
}

func (s *concurrentSet[T]) Union(other Set[T]) Set[T] {
	o, _ := other.(*concurrentSet[T])

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	union, _ := s.ncs.Union(o.ncs).(*noconcurrentSet[T])

	return &concurrentSet[T]{
		ncs: union,
	}
}

func (s *concurrentSet[T]) Intersection(other Set[T]) Set[T] {
	o, _ := other.(*concurrentSet[T])

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	intersection, _ := s.ncs.Intersection(o.ncs).(*noconcurrentSet[T])

	return &concurrentSet[T]{
		ncs: intersection,
	}
}

func (s *concurrentSet[T]) Difference(other Set[T]) Set[T] {
	o, _ := other.(*concurrentSet[T])

	s.RLock()
	o.RLock()
	defer s.RUnlock()
	defer o.RUnlock()

	difference, _ := s.ncs.Difference(o.ncs).(*noconcurrentSet[T])

	return &concurrentSet[T]{
		ncs: difference,
	}
}
