package gosl

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// RandomString generates a random string with given size using built-in
// "crypto/rand" and "encoding/hex" packages.
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
//		s, err := gosl.RandomString(8)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(s)
//	}
func RandomString(size int) (string, error) {
	if size <= 0 {
		return "", errors.New("can't generate random string with zero or negative size")
	}

	bufferSize := size
	if size%2 == 0 && size >= 4 {
		bufferSize = size / 2 // hack to reduce buffer size and B/op
	}

	b := make([]byte, bufferSize)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.New("can't read given buffer")
	}

	return hex.EncodeToString(b)[:size], nil
}
