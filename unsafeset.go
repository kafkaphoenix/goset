package goset

type unsafeSet[T comparable] map[T]struct{}

// Interface guard to assert concrete type:unsafeSet adheres to Set interface.
var _ Set[string] = (*unsafeSet[string])(nil)

func newUnsafeSet[T comparable](size int) *unsafeSet[T] {
	s := make(unsafeSet[T], size)
	return &s
}

func (s *unsafeSet[T]) Add(v T) {
	(*s)[v] = struct{}{}
}

func (s *unsafeSet[T]) Remove(v T) {
	delete(*s, v)
}

func (s *unsafeSet[T]) Contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}

func (s *unsafeSet[T]) Size() int {
	return len(*s)
}

func (s *unsafeSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *unsafeSet[T]) IsEqual(other Set[T]) bool {
	o, _ := other.(*unsafeSet[T])

	if s.Size() != o.Size() {
		return false
	}

	for key := range *s {
		if !o.Contains(key) {
			return false
		}
	}

	return true
}

func (s *unsafeSet[T]) Clear() {
	*s = make(map[T]struct{})
}

func (s *unsafeSet[T]) ToSlice() []T {
	list := make([]T, 0, s.Size())
	for key := range *s {
		list = append(list, key)
	}

	return list
}

func (s *unsafeSet[T]) Clone() Set[T] {
	clone := newUnsafeSet[T](s.Size())
	for key := range *s {
		clone.Add(key)
	}

	return clone
}

func (s *unsafeSet[T]) Union(other Set[T]) Set[T] {
	o, _ := other.(*unsafeSet[T])

	n := max(o.Size(), s.Size())
	union := newUnsafeSet[T](n)

	for key := range *s {
		union.Add(key)
	}

	for key := range *o {
		union.Add(key)
	}

	return union
}

func (s *unsafeSet[T]) Intersection(other Set[T]) Set[T] {
	o, _ := other.(*unsafeSet[T])

	intersect := NewSet[T]()

	if s.Size() < other.Size() {
		for key := range *s {
			if other.Contains(key) {
				intersect.Add(key)
			}
		}
	} else {
		for key := range *o {
			if s.Contains(key) {
				intersect.Add(key)
			}
		}
	}

	return intersect
}

func (s *unsafeSet[T]) Difference(other Set[T]) Set[T] {
	o, _ := other.(*unsafeSet[T])

	diff := NewSet[T]()

	for key := range *s {
		if !o.Contains(key) {
			diff.Add(key)
		}
	}

	return diff
}
