package gosl

import (
	"os"
	"strings"
)

// ContainsCaseInsensitive reports if substr is within s string using built-in
// "strings" package with strings.Contains. Case-insensitive for input values by
// default.
//
// If s and/or substr have a zero-value, returns false for bool.
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
// If s has a zero-value, returns false for bool.
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
// If m has a zero-value, returns false for bool.
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
func ContainsInMap[T any, K comparable](m map[K]T, key K) bool {
	if m == nil {
		return false
	}

	if _, exists := m[key]; exists {
		return true
	}

	return false
}

// IsFileExist reports whether a file exists on the specified path.
//
// If path has a zero-value or is dir, returns false for bool.
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
//		p := filepath.Clean("~/Downloads/file.csv")
//
//		b := gosl.IsFileExist(p)
//
//		fmt.Println(b)
//	}
func IsFileExist(path string) bool {
	// Check, if the specified path has a zero-value.
	if path == "" {
		return false
	}

	// Get stat of the specified path or error.
	file, err := os.Stat(path)

	// Check, if file is not dir.
	if file.IsDir() {
		return false
	}

	return err == nil || !os.IsNotExist(err)
}

// IsDirExist reports whether a dir exists on the specified path.
//
// If path has a zero-value or is file, returns false for bool.
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
//		p := filepath.Clean("~/Downloads/my-folder")
//
//		b := gosl.IsDirExist(p)
//
//		fmt.Println(b)
//	}
func IsDirExist(path string) bool {
	// Check, if the specified path has a zero-value.
	if path == "" {
		return false
	}

	// Get stat of the specified path or error.
	dir, err := os.Stat(path)

	// Check, if dir is not file.
	if !dir.IsDir() {
		return false
	}

	return err == nil || !os.IsNotExist(err)
}
