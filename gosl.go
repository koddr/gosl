// Package gosl provides a snippet collection for working with routine
// operations in your Go programs with a super user-friendly API and the most
// efficient performance.
package gosl

import "github.com/charmbracelet/lipgloss"

// Utility represents struct for a regular function.
type Utility struct{}

// GenericUtility represents struct with T any and K comparable types for a
// generic function.
type GenericUtility[T any, K comparable] struct{}

// Concat concatenate strings using the built-in copy and "unsafe" package with
// unsafe.String function.
//
// If s has no elements returns zero-value for a string.
func (u *Utility) Concat(s ...string) string {
	return Concat(s...)
}

// ContainsCaseInsensitive reports if substr is within s string using built-in
// "strings" package with strings.Contains. Case-insensitive for input values by
// default.
//
// If s and/or substr have an "" (empty) value returns false for a bool.
func (u *Utility) ContainsCaseInsensitive(s, substr string) bool {
	return ContainsCaseInsensitive(s, substr)
}

// RandomString generates a random string with a given size using built-in
// "crypto/rand" and "encoding/hex" packages.
//
// If err != nil returns zero-value for a string and error.
func (u *Utility) RandomString(size int) (string, error) {
	return RandomString(size)
}

// RenderStyled render a styled string with a given lipgloss.Style template
// using "charmbracelet/lipgloss" package.
//
// If s have an "" (empty) value returns zero-value for a string.
func (u *Utility) RenderStyled(s string, template lipgloss.Style) string {
	return RenderStyled(s, template)
}

// ToBytes converts string to byte slice using the built-in "unsafe" package
// with unsafe.Slice function.
//
// If err != nil returns zero-value for a byte slice and error.
func (u *Utility) ToBytes(s string) ([]byte, error) {
	return ToBytes(s)
}

// ToString converts byte slice to string using the built-in "unsafe" package
// with unsafe.String function.
//
// If err != nil returns zero-value for a string and error.
func (u *Utility) ToString(b []byte) (string, error) {
	return ToString(b)
}

// ContainsInSlice reports if value T is within slice []T.
//
// If s have a zero-value returns false for a bool.
func (g *GenericUtility[T, K]) ContainsInSlice(s []K, value K) bool {
	return ContainsInSlice(s, value)
}

// ContainsInMap reports if key T is within map[T]K.
//
// If m have a zero-value returns false for a bool.
func (g *GenericUtility[T, K]) ContainsInMap(m map[K]T, key K) bool {
	return ContainsInMap(m, key)
}

// Equals compares two values of type K, returns true if they are equal.
func (g *GenericUtility[T, K]) Equals(value1, value2 K) bool {
	return Equals(value1, value2)
}

// NotEquals compares two values of type K, returns true if they are not equal.
func (g *GenericUtility[T, K]) NotEquals(value1, value2 K) bool {
	return NotEquals(value1, value2)
}

// Marshal converts struct *T to JSON data (byte slice) using jsoniter.Marshal
// with a default configuration. A 100% compatible drop-in replacement of
// "encoding/json" standard lib.
//
// If err != nil returns zero-value for a byte slice and error.
func (g *GenericUtility[T, K]) Marshal(model *T) ([]byte, error) {
	return Marshal(model)
}

// Unmarshal converts JSON data (byte slice) to struct *T using
// jsoniter.Unmarshal with a default configuration. A 100% compatible drop-in
// replacement of "encoding/json" standard lib.
//
// If err != nil returns zero-value for a struct and error.
func (g *GenericUtility[T, K]) Unmarshal(data []byte, model *T) (*T, error) {
	return Unmarshal(data, model)
}
