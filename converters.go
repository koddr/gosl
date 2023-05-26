package gosl

import (
	"errors"
	"unsafe"
)

// ToBytes converts string to byte slice using the built-in "unsafe" package
// with unsafe.Slice function.
//
// If err != nil returns zero-value for a byte slice and error.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//		"log"
//
//		"github.com/koddr/gosl"
//	)
//
//	func main() {
//		s := "this is my string"
//
//		b, err := gosl.ToBytes(s)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(string(b))
//	}
func ToBytes(s string) ([]byte, error) {
	if s == "" {
		return nil, errors.New("can't convert empty string to byte slice")
	}

	return unsafe.Slice(unsafe.StringData(s), len(s)), nil
}

// ToString converts byte slice to string using the built-in "unsafe" package
// with unsafe.String function.
//
// If err != nil returns zero-value for a string and error.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//		"log"
//
//		"github.com/koddr/gosl"
//	)
//
//	func main() {
//		b := []byte("this is my string")
//
//		s, err := gosl.ToString(b)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(s)
//	}
func ToString(b []byte) (string, error) {
	if b == nil {
		return "", errors.New("can't convert nil byte slice to string")
	}

	return unsafe.String(unsafe.SliceData(b), len(b)), nil
}
