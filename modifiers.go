package gosl

import (
	"reflect"
)

// ModifyByValue modify an unknown key in the given map[string]any by it value.
// Supports nested maps, but only if their type is map[string]any.
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
//		m := map[string]any{"order": map[string]any{"total_cost": 100}}
//		foundValue := 100
//		newValue := 250
//
//		isFound, result := gosl.ModifyByValue(m, foundValue, newValue)
//
//		fmt.Println(isFound, result)
//	}
func ModifyByValue(m map[string]any, foundValue, newValue any) (foundKey bool, results map[string]any) {
	// Check, if the given map is empty.
	if m == nil {
		return foundKey, nil
	}

	// Loop for all keys in the given map.
	for key, value := range m {
		// Check map by reflect.
		if reflect.DeepEqual(value, foundValue) {
			// Modify a key of the given map.
			m[key] = newValue
			foundKey = true // switch the temp variable
		} else if mv, ok := value.(map[string]any); ok {
			// Run recurrent function.
			isFound, modified := ModifyByValue(mv, foundValue, newValue)

			// Check, if key is found.
			if isFound {
				// Modify a key of the given map.
				m[key] = modified
				foundKey = true // switch the temp variable
			}
		}
	}

	return foundKey, m
}
