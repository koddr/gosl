package gosl

import (
	"errors"
	"unsafe"
)

// ToBytes function to convert string to byte slice. Using built-in "unsafe"
// package with unsafe.Slice function.
//
// If err != nil returns nil value for a byte slice and error.
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
func ToBytes(str string) ([]byte, error) {
	if str == "" {
		return nil, errors.New("can't convert empty string to byte slice")
	}
	
	return unsafe.Slice(unsafe.StringData(str), len(str)), nil
}

// ToString function to convert byte slice to string. Using built-in "unsafe"
// package with unsafe.String function.
//
// If err != nil returns "" (empty) value for a string and error.
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
func ToString(byteSlice []byte) (string, error) {
	if byteSlice == nil {
		return "", errors.New("can't convert nil byte slice to string")
	}
	
	return unsafe.String(unsafe.SliceData(byteSlice), len(byteSlice)), nil
}
