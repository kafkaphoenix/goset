//go:build unit

package goset_test

import (
	"github.com/kafkaphoenix/goset"
	"slices"
	"testing"
)

func TestNewSet_OK(t *testing.T) {
	tests := map[string]struct {
		values   []string
		expected goset.Set[string]
	}{
		"empty": {
			values:   []string{},
			expected: goset.NewSet[string](),
		},
		"single": {
			values:   []string{"a"},
			expected: goset.NewSet("a"),
		},
		"multiple": {
			values:   []string{"a", "b", "c"},
			expected: goset.NewSet("a", "b", "c"),
		},
		"duplicates": {
			values:   []string{"a", "b", "a", "c"},
			expected: goset.NewSet("a", "b", "c"),
		},
		"unordered": {
			values:   []string{"a", "b", "c"},
			expected: goset.NewSet("c", "a", "b"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			set := goset.NewSet(test.values...)
			if !set.IsEqual(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected.ToSlice(), set.ToSlice())
			}
		})
	}
}

func runGenericSetTest[T comparable](t *testing.T, values []T, expected goset.Set[T]) {
	t.Helper()
	set := goset.NewSet(values...)
	if !set.IsEqual(expected) {
		t.Fatalf("expected %v, got %v", expected.ToSlice(), set.ToSlice())
	}
}

func TestGeneric_OK(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		runGenericSetTest(t, []string{"a", "b", "c"}, goset.NewSet("a", "b", "c"))
	})
	t.Run("int", func(t *testing.T) {
		runGenericSetTest(t, []int{1, 2, 3}, goset.NewSet(1, 2, 3))
	})
	t.Run("float", func(t *testing.T) {
		runGenericSetTest(t, []float64{1.1, 2.2, 3.3}, goset.NewSet(1.1, 2.2, 3.3))
	})
	t.Run("struct", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		points := []Point{{1, 2}, {3, 4}, {5, 6}}
		expected := goset.NewSet(points...)
		runGenericSetTest(t, points, expected)
	})
	t.Run("bool", func(t *testing.T) {
		runGenericSetTest(t, []bool{true, false}, goset.NewSet(true, false))
	})
	t.Run("byte", func(t *testing.T) {
		runGenericSetTest(t, []byte{1, 2, 3}, goset.NewSet(byte(1), byte(2), byte(3)))
	})
	t.Run("rune", func(t *testing.T) {
		runGenericSetTest(t, []rune{'a', 'b', 'c'}, goset.NewSet('a', 'b', 'c'))
	})
	t.Run("complex", func(t *testing.T) {
		runGenericSetTest(t, []complex128{1 + 2i, 3 + 4i}, goset.NewSet(complex(1, 2), complex(3, 4)))
	})
}

func TestAdd_OK(t *testing.T) {
	tests := map[string]struct {
		initial  goset.Set[string]
		add      string
		expected goset.Set[string]
	}{
		"add_unique": {
			initial:  goset.NewSet("a", "b"),
			add:      "c",
			expected: goset.NewSet("a", "b", "c"),
		},
		"add_duplicate": {
			initial:  goset.NewSet("a", "b"),
			add:      "a",
			expected: goset.NewSet("a", "b"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.initial.Add(test.add)
			if !test.initial.IsEqual(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected.ToSlice(), test.initial.ToSlice())
			}
		})
	}
}

func TestRemove_OK(t *testing.T) {
	tests := map[string]struct {
		initial  goset.Set[string]
		remove   string
		expected goset.Set[string]
	}{
		"remove_existing": {
			initial:  goset.NewSet("a", "b", "c"),
			remove:   "b",
			expected: goset.NewSet("a", "c"),
		},
		"remove_non_existing": {
			initial:  goset.NewSet("a", "b"),
			remove:   "c",
			expected: goset.NewSet("a", "b"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.initial.Remove(test.remove)
			if !test.initial.IsEqual(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected.ToSlice(), test.initial.ToSlice())
			}
		})
	}
}

func TestContains_OK(t *testing.T) {
	tests := map[string]struct {
		initial  goset.Set[string]
		contains string
		expected bool
	}{
		"contains_existing": {
			initial:  goset.NewSet("a", "b", "c"),
			contains: "b",
			expected: true,
		},
		"contains_non_existing": {
			initial:  goset.NewSet("a", "b"),
			contains: "c",
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.initial.Contains(test.contains) != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, test.initial.Contains(test.contains))
			}
		})
	}
}

func TestSize_OK(t *testing.T) {
	tests := map[string]struct {
		initial  goset.Set[string]
		expected int
	}{
		"empty": {
			initial:  goset.NewSet[string](),
			expected: 0,
		},
		"single": {
			initial:  goset.NewSet("a"),
			expected: 1,
		},
		"multiple": {
			initial:  goset.NewSet("a", "b", "c"),
			expected: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.initial.Size() != test.expected {
				t.Fatalf("expected %d, got %d", test.expected, test.initial.Size())
			}
		})
	}
}

func TestIsEmpty_OK(t *testing.T) {
	tests := map[string]struct {
		initial  goset.Set[string]
		expected bool
	}{
		"empty": {
			initial:  goset.NewSet[string](),
			expected: true,
		},
		"non_empty": {
			initial:  goset.NewSet("a"),
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.initial.IsEmpty() != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, test.initial.IsEmpty())
			}
		})
	}
}

func TestIsEqual_OK(t *testing.T) {
	tests := map[string]struct {
		set1     goset.Set[string]
		set2     goset.Set[string]
		expected bool
	}{
		"equal": {
			set1:     goset.NewSet("a", "b", "c"),
			set2:     goset.NewSet("a", "b", "c"),
			expected: true,
		},
		"equal_unordered": {
			set1:     goset.NewSet("a", "b", "c"),
			set2:     goset.NewSet("c", "a", "b"),
			expected: true,
		},
		"not_equal_different_size": {
			set1:     goset.NewSet("a", "b"),
			set2:     goset.NewSet("a", "b", "c"),
			expected: false,
		},
		"not_equal_different_elements": {
			set1:     goset.NewSet("a", "b"),
			set2:     goset.NewSet("a", "c"),
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.set1.IsEqual(test.set2) != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, test.set1.IsEqual(test.set2))
			}
		})
	}
}

func TestClear_OK(t *testing.T) {
	tests := map[string]struct {
		initial  goset.Set[string]
		expected goset.Set[string]
	}{
		"clear_non_empty": {
			initial:  goset.NewSet("a", "b", "c"),
			expected: goset.NewSet[string](),
		},
		"clear_empty": {
			initial:  goset.NewSet[string](),
			expected: goset.NewSet[string](),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.initial.Clear()
			if !test.initial.IsEqual(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected.ToSlice(), test.initial.ToSlice())
			}
		})
	}
}

func TestToSlice_OK(t *testing.T) {
	tests := map[string]struct {
		initial  goset.Set[string]
		expected []string
	}{
		"empty": {
			initial:  goset.NewSet[string](),
			expected: []string{},
		},
		"single": {
			initial:  goset.NewSet("a"),
			expected: []string{"a"},
		},
		"multiple": {
			initial:  goset.NewSet("a", "b", "c"),
			expected: []string{"a", "b", "c"},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			slice := test.initial.ToSlice()
			if name == "multiple" {
				// This is necessary because the order of elements in a set is not guaranteed
				slices.Sort(slice)
				slices.Sort(test.expected)
			}
			if len(slice) != len(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected, slice)
			}
			for i, v := range slice {
				if v != test.expected[i] {
					t.Fatalf("expected %v, got %v", test.expected[i], v)
				}
			}
		})
	}
}

func TestClone_OK(t *testing.T) {
	tests := map[string]struct {
		initial  goset.Set[string]
		expected goset.Set[string]
	}{
		"empty": {
			initial:  goset.NewSet[string](),
			expected: goset.NewSet[string](),
		},
		"single": {
			initial:  goset.NewSet("a"),
			expected: goset.NewSet("a"),
		},
		"multiple": {
			initial:  goset.NewSet("a", "b", "c"),
			expected: goset.NewSet("a", "b", "c"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			clone := test.initial.Clone()
			if !clone.IsEqual(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected.ToSlice(), clone.ToSlice())
			}
		})
	}
}

func TestUnion_OK(t *testing.T) {
	tests := map[string]struct {
		set1     goset.Set[string]
		set2     goset.Set[string]
		expected goset.Set[string]
	}{
		"empty_union": {
			set1:     goset.NewSet[string](),
			set2:     goset.NewSet[string](),
			expected: goset.NewSet[string](),
		},
		"non_empty_union": {
			set1:     goset.NewSet("a", "b"),
			set2:     goset.NewSet("b", "c"),
			expected: goset.NewSet("a", "b", "c"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			union := test.set1.Union(test.set2)
			if !union.IsEqual(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected.ToSlice(), union.ToSlice())
			}
		})
	}
}

func TestIntersection_OK(t *testing.T) {
	tests := map[string]struct {
		set1     goset.Set[string]
		set2     goset.Set[string]
		expected goset.Set[string]
	}{
		"empty_intersection": {
			set1:     goset.NewSet[string](),
			set2:     goset.NewSet[string](),
			expected: goset.NewSet[string](),
		},
		"non_empty_intersection": {
			set1:     goset.NewSet("a", "b"),
			set2:     goset.NewSet("b", "c"),
			expected: goset.NewSet("b"),
		},
		"second_biggest": {
			set1:     goset.NewSet("b", "c"),
			set2:     goset.NewSet("b", "c", "d"),
			expected: goset.NewSet("b", "c"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			intersection := test.set1.Intersection(test.set2)
			if !intersection.IsEqual(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected.ToSlice(), intersection.ToSlice())
			}
		})
	}
}

func TestDifference_OK(t *testing.T) {
	tests := map[string]struct {
		set1     goset.Set[string]
		set2     goset.Set[string]
		expected goset.Set[string]
	}{
		"empty_difference": {
			set1:     goset.NewSet[string](),
			set2:     goset.NewSet[string](),
			expected: goset.NewSet[string](),
		},
		"non_empty_difference": {
			set1:     goset.NewSet("a", "b"),
			set2:     goset.NewSet("b", "c"),
			expected: goset.NewSet("a"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			difference := test.set1.Difference(test.set2)
			if !difference.IsEqual(test.expected) {
				t.Fatalf("expected %v, got %v", test.expected.ToSlice(), difference.ToSlice())
			}
		})
	}
}
