//go:build unit

package goset_test

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"

	"github.com/kafkaphoenix/goset"
)

const N = 1000

func TestConcurrent_Add_OK(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := goset.NewSafeSet[int]()
	values := rand.Perm(N)

	var wg sync.WaitGroup
	wg.Add(len(values))
	for i := range len(values) {
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(i)
	}
	wg.Wait()
	for _, i := range values {
		if !s.Contains(i) {
			t.Errorf("Set is missing expected element: %v", i)
		}
	}
}

func TestConcurrent_Remove_OK(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := goset.NewSafeSet[int]()
	values := rand.Perm(N)
	for _, v := range values {
		s.Add(v)
	}

	var wg sync.WaitGroup
	wg.Add(len(values))
	for _, v := range values {
		go func(i int) {
			defer wg.Done()
			s.Remove(i)
		}(v)
	}
	wg.Wait()

	if s.Size() != 0 {
		t.Errorf("Set is not empty after removing elements")
	}
}

func TestConcurrent_Contains_OK(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := goset.NewSafeSet[int]()
	values := rand.Perm(N)
	for _, v := range values {
		s.Add(v)
	}

	var wg sync.WaitGroup
	for _, v := range values {
		number := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !s.Contains(number) {
				t.Errorf("Set is missing expected element: %v", number)
			}
		}()
	}
	wg.Wait()
}

func TestConcurrent_Size_OK(t *testing.T) {
	runtime.GOMAXPROCS(2)

	s := goset.NewSafeSet[int]()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		size := s.Size()
		for range N {
			newSize := s.Size()
			if newSize < size {
				t.Errorf("Size changed unexpectedly from %d to %d", size, newSize)
			}
		}
	}()

	for range N {
		s.Add(rand.Int())
	}
	wg.Wait()
}

// func TestConcurrent_IsEmpty_OK(t *testing.T) {
// 	runtime.GOMAXPROCS(2)

// 	s := goset.NewSafeSet[int]()
// 	var wg sync.WaitGroup

// 	// concurrent writer
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for range N {
// 			val := rand.Intn(N)
// 			s.Add(val)
// 			s.Remove(val)
// 		}
// 	}()

// 	// concurrent reader
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for range N {
// 			cs := s.(*goset.TestSafeSet[int])
// 			cs.RLock()
// 			// deadlock and without this empty and size fail
// 			// because each has a lock
// 			if cs.IsEmpty() {
// 				if cs.Size() != 0 {
// 					t.Errorf("Set is empty but size is not 0")
// 				}
// 			}
// 			cs.RUnlock()
// 		}
// 	}()

// 	wg.Wait()
// }

// func TestConcurrent_IsEqual_OK(t *testing.T) {
// 	runtime.GOMAXPROCS(2)

// 	s1 := goset.NewSafeSet[int]()
// 	s2 := goset.NewSafeSet[int]()

// 	values := rand.Perm(N)
// 	for _, v := range values {
// 		s1.Add(v)
// 		s2.Add(v)
// 	}

// 	var wg sync.WaitGroup
// 	for range N {
// 		wg.Add(1)
// 		go func() {
//			defer wg.Done()
// 			s1.IsEqual(s2)
// 		}()
// 	}
// 	wg.Wait()
// }

// func TestConcurrent_IsEqual_OK(t *testing.T) {
// 	runtime.GOMAXPROCS(2)

// 	s1 := goset.NewSafeSet[int]()
// 	s2 := goset.NewSafeSet[int]()

// 	values := rand.Perm(N)
// 	for _, v := range values {
// 		s1.Add(v)
// 		s2.Add(v)
// 	}

// 	var wg sync.WaitGroup

// 	// Concurrently mutate s
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < N; i++ {
// 			s1.Add(rand.Intn(2 * N))
// 			s1.Remove(rand.Intn(2 * N))
// 		}
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for range N {
// 			_ = s1.IsEqual(s2)
// 		}
// 	}()

// 	wg.Wait()
// }
