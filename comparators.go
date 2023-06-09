package gosl

// Equals compares two values of type T, returns true if they are equal.
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
//		s1 := "hello"
//		s2 := "hello"
//
//		b := gosl.Equals(s1, s2)
//
//		fmt.Println(b)
//	}
func Equals[T comparable](value1, value2 T) bool {
	return value1 == value2
}

// NotEquals compares two values of type T, returns true if they are not equal.
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
//		s1 := "hello"
//		s2 := "world"
//
//		b := gosl.NotEquals(s1, s2)
//
//		fmt.Println(b)
//	}
func NotEquals[T comparable](value1, value2 T) bool {
	return value1 != value2
}
