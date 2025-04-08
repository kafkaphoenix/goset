package main

import (
	"fmt"
	"sync"

	set "github.com/kafkaphoenix/goset"
)

func nonConcurrentDemo() {
	fmt.Println("///////////Non-concurrent set demo")
	// Create a new set

	s := set.NewSet("a", "b")

	// Add elements to the set
	s.Add("c")
	s.Add("d")

	fmt.Printf("Set: %s\n", s.ToSlice()) // [a b c d]

	// Add duplicate elements
	s.Add("a")
	fmt.Println("After adding duplicates:", s.ToSlice()) // [a b c d]

	// Check if an element is in the set
	fmt.Println("Contains 'a':", s.Contains("a")) // true
	fmt.Println("Contains 'e':", s.Contains("e")) // false

	// Remove an element from the set
	s.Remove("b")
	fmt.Println("After removing 'b':", s.ToSlice()) // [a c d]
	s.Remove("e")                                   // no error
	fmt.Println("After removing 'e':", s.ToSlice()) // [a c d]

	// Check if the set is empty
	fmt.Println("Is empty:", s.IsEmpty()) // false

	// Check if the set is equal to another set
	otherSet := set.NewSet("a", "c", "d")
	fmt.Printf("Is equal to set %s: %v\n", otherSet.ToSlice(), s.IsEqual(otherSet)) // true

	// Get the size of the set
	fmt.Println("Size:", s.Size()) // 3

	// Clear the set
	s.Clear()
	fmt.Println("After clearing:", s.ToSlice()) // []
	fmt.Println("Is empty:", s.IsEmpty())       // true
}

func concurrentDemo() {
	fmt.Println("///////////Concurrent safe set demo")
	// Create a new concurrent safe set
	cs := set.NewSafeSet[int]()

	const numGoroutines = 5

	const opsPerGoroutine = 5

	var wg sync.WaitGroup

	// Start multiple goroutines to add and remove elements
	fmt.Println("Adding and removing elements in concurrent safe set...")
	fmt.Printf("Number of goroutines: %d, Operations per goroutine: %d\n", numGoroutines, opsPerGoroutine)
	fmt.Println("Removing even numbers...")
	fmt.Println("Adding odd numbers...")

	for i := range numGoroutines {
		wg.Add(1)

		go func(base int) {
			defer wg.Done()

			for j := 0; j < opsPerGoroutine; j++ {
				v := base*opsPerGoroutine + j
				cs.Add(v)

				if v%2 == 0 {
					cs.Remove(v) // remove even numbers
				}
			}
		}(i)
	}

	wg.Wait()

	// Validate that only odd numbers remain
	slice := cs.ToSlice()
	for _, v := range slice {
		if v%2 == 0 {
			fmt.Printf("Example failed...found even number %d in the set\n", v)
		}
	}

	fmt.Printf("Example completed...found %d odd numbers in the set\n", len(slice))
	fmt.Printf("Set: %v\n", slice) // should contain only odd numbers
}

func setOperations() {
	fmt.Println("///////////Set operations demo")

	a := set.NewSet(1, 2, 3)
	b := set.NewSet(3, 4, 5)

	fmt.Println("Set A:", a.ToSlice()) // [1 2 3]
	fmt.Println("Set B:", b.ToSlice()) // [3 4 5]

	// Perform set operations
	union := a.Union(b)
	inter := a.Intersection(b)
	diff := a.Difference(b)

	fmt.Println("Union:", union.ToSlice())        // [1 2 3 4 5]
	fmt.Println("Intersection:", inter.ToSlice()) // [3]
	fmt.Println("Difference:", diff.ToSlice())    // [1 2]

	clone := a.Clone()
	fmt.Println("Original:", a.ToSlice())            // [1 2 3]
	fmt.Println("Clone:", clone.ToSlice())           // [1 2 3]
	fmt.Println("Clone is equal?", clone.IsEqual(a)) // true
}

func main() {
	nonConcurrentDemo()
	concurrentDemo()
	setOperations()
}
