package gosl

import "unsafe"

// Concat concatenate strings using the built-in copy and "unsafe" package with
// unsafe.String function.
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
//		s1 := "this "
//		s2 := "is "
//		s3 := "my string"
//
//		s := gosl.Concat(s1, s2, s3)
//
//		fmt.Println(s)
//	}
func Concat(s ...string) string {
	if len(s) == 0 {
		return ""
	}

	n := 0
	for i := 0; i < len(s); i++ {
		n += len(s[i])
	}

	b := make([]byte, n, n)

	idx := 0
	for i := 0; i < len(s); i++ {
		copy(b[idx:], s[i])
		idx += len(s[i])
	}

	return unsafe.String(unsafe.SliceData(b), n)
}
