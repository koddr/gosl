package gosl

import "strings"

// ContainsCaseInsensitive reports if substr is within s string using built-in
// "strings" package with strings.Contains. Case-insensitive for input values by
// default.
//
// If s and/or substr have a zero value returns false value for a bool.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/koddr/gosl"
//	)
//
//	func main() {
//		s := "this is my string"
//		substr := "my"
//
//		b := gosl.ContainsCaseInsensitive(s, substr)
//
//		fmt.Println(b)
//	}
func ContainsCaseInsensitive(s, substr string) bool {
	if s == "" || substr == "" {
		return false
	}

	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// ContainsInSlice reports if value T is within slice []T.
//
// If s have a zero value returns false value for a bool.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/koddr/gosl"
//	)
//
//	func main() {
//		s := []string{"one", "two", "three"}
//		value := "two"
//
//		b := gosl.ContainsInSlice(s, value)
//
//		fmt.Println(b)
//	}
func ContainsInSlice[T comparable](s []T, value T) bool {
	if s == nil {
		return false
	}

	for _, elem := range s {
		if elem == value {
			return true
		}
	}

	return false
}

// ContainsInMap reports if key T is within map[T]K.
//
// If m have a zero value returns false value for a bool.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/koddr/gosl"
//	)
//
//	func main() {
//		m := map[string]int{"one": 1, "two": 2, "three": 3}
//		key := "two"
//
//		b := gosl.ContainsInMap(m, key)
//
//		fmt.Println(b)
//	}
func ContainsInMap[T comparable, K any](m map[T]K, key T) bool {
	if m == nil {
		return false
	}

	if _, exists := m[key]; exists {
		return true
	}

	return false
}
