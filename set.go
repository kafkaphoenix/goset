package goset

// Set defines the behavior of a generic Set.
type Set[T comparable] interface {
	// Add inserts a value into the set.
	Add(value T)
	// Remove deletes a value from the set.
	Remove(value T)
	// Contains checks whether a value is in the set.
	Contains(value T) bool
	// Size returns the number of elements in the set.
	Size() int
	// IsEmpty returns true if the set has no elements.
	IsEmpty() bool
	// IsEqual checks if two sets are equal.
	IsEqual(other Set[T]) bool
	// Clear removes all elements from the set.
	Clear()
	// ToSlice returns a slice of all elements in the set.
	// WARNING Order is not guaranteed.
	ToSlice() []T
	// Clone creates a deep copy of the set.
	Clone() Set[T]
	// Union returns a new set with elements from both sets.
	Union(other Set[T]) Set[T]
	// Intersection returns a new set with elements common to both sets.
	Intersection(other Set[T]) Set[T]
	// Difference returns a new set with elements in the current set but not in the other.
	Difference(other Set[T]) Set[T]
}

// NewSet creates a set with the given elements.
// It is not safe for concurrent use.
// Use NewSafeSet for concurrent use.
func NewSet[T comparable](vs ...T) Set[T] {
	s := newUnsafeSet[T](len(vs))
	for _, v := range vs {
		s.Add(v)
	}

	return s
}

// NewSafeSet creates a concurrent safe set with the given elements.
func NewSafeSet[T comparable](vs ...T) Set[T] {
	s := newSafeSet[T](len(vs))
	for _, v := range vs {
		s.Add(v)
	}

	return s
}
