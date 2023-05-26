package gosl

import jsoniter "github.com/json-iterator/go"

// Marshal converts struct *T to JSON data (byte slice) using jsoniter.Marshal
// with a default configuration. A 100% compatible drop-in replacement of
// "encoding/json" standard lib.
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
//	type user struct {
//		ID   int    `json:"id"`
//		Name string `json:"name"`
//	}
//
//	func main() {
//		u := &user{ID: 1, Name: "Viktor"}
//
//		json, err := gosl.Marshal(u)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(string(json))
//	}
func Marshal[T any](model *T) ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&model)
}

// Unmarshal converts JSON data (byte slice) to struct *T using
// jsoniter.Unmarshal with a default configuration. A 100% compatible drop-in
// replacement of "encoding/json" standard lib.
//
// If err != nil returns zero-value for a struct and error.
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
//	type user struct {
//		ID   int    `json:"id"`
//		Name string `json:"name"`
//	}
//
//	func main() {
//		json := []byte(`{"id":1,"name":"Viktor"}`)
//		model := &user{}
//
//		u, err := gosl.Unmarshal(json, model)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(u)
//	}
func Unmarshal[T any](data []byte, model *T) (*T, error) {
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &model); err != nil {
		return nil, err
	}

	return model, nil
}
