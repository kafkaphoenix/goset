# Goset

## Description
Goset is a Go library that provides a simple and efficient way to work with [generic](https://go.dev/blog/intro-generics) sets. It allows you to create, manipulate, and perform operations on sets. It offers both a concurrent and no concurrent implementation, allowing you to choose the one that best fits your needs.

## Features
- Create sets of [comparable](https://go.dev/blog/comparable) elements.
- Perform set operations such as union, intersection, and difference.

## Example
```go
package main

import (
    "fmt"
    "github.com/yourusername/goset"
)

func main() {
    // Create a new set
    s := goset.NewSet[int]()
    
    // Add elements to the set
    s.Add(1)
    s.Add(2)
    s.Add(3)
    
    // Check if an element is in the set
    fmt.Println(s.Contains(2)) // true
    
    // Remove an element from the set
    s.Remove(2)
    
    // Get the size of the set
    fmt.Println(s.Size()) // 2
    
    // Iterate over the elements in the set
    for _, v := range s.Values() {
        fmt.Println(v) // 1, 3
    }
}
```