package goset

type noconcurrentSet[T comparable] map[T]struct{}

// Interface guard to assert concrete type:noconcurrentSet adheres to Set interface.
var _ Set[string] = (*noconcurrentSet[string])(nil)

func newNoConcurrentSet[T comparable](size int) *noconcurrentSet[T] {
	s := make(noconcurrentSet[T], size)
	return &s
}

func (s *noconcurrentSet[T]) Add(v T) {
	(*s)[v] = struct{}{}
}

func (s *noconcurrentSet[T]) Remove(v T) {
	delete(*s, v)
}

func (s *noconcurrentSet[T]) Contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}

func (s *noconcurrentSet[T]) Size() int {
	return len(*s)
}

func (s *noconcurrentSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *noconcurrentSet[T]) IsEqual(other Set[T]) bool {
	o, _ := other.(*noconcurrentSet[T])

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

func (s *noconcurrentSet[T]) Clear() {
	*s = make(map[T]struct{})
}

func (s *noconcurrentSet[T]) ToSlice() []T {
	list := make([]T, 0, s.Size())
	for key := range *s {
		list = append(list, key)
	}

	return list
}

func (s *noconcurrentSet[T]) Clone() Set[T] {
	clone := newNoConcurrentSet[T](s.Size())
	for key := range *s {
		clone.Add(key)
	}

	return clone
}

func (s *noconcurrentSet[T]) Union(other Set[T]) Set[T] {
	o, _ := other.(*noconcurrentSet[T])

	n := max(o.Size(), s.Size())
	union := newNoConcurrentSet[T](n)

	for key := range *s {
		union.Add(key)
	}

	for key := range *o {
		union.Add(key)
	}

	return union
}

func (s *noconcurrentSet[T]) Intersection(other Set[T]) Set[T] {
	o, _ := other.(*noconcurrentSet[T])

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

func (s *noconcurrentSet[T]) Difference(other Set[T]) Set[T] {
	o, _ := other.(*noconcurrentSet[T])

	diff := NewSet[T]()

	for key := range *s {
		if !o.Contains(key) {
			diff.Add(key)
		}
	}

	return diff
}
