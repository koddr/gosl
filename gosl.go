// Package gosl provides snippets collection for working with routine operations
// in your Go programs with a super user-friendly API and the most efficient
// performance.
package gosl

type Utility struct{}
type GenericUtility[T any, K comparable] struct{}

func (u *Utility) ToBytes(s string) ([]byte, error) {
	return ToBytes(s)
}

func (u *Utility) ToString(b []byte) (string, error) {
	return ToString(b)
}

func (u *Utility) RandomString(size int) (string, error) {
	return RandomString(size)
}

func (u *Utility) ContainsCaseInsensitive(s, substr string) bool {
	return ContainsCaseInsensitive(s, substr)
}

func (g *GenericUtility[T, K]) ContainsInSlice(s []K, value K) bool {
	return ContainsInSlice(s, value)
}

func (g *GenericUtility[T, K]) ContainsInMap(m map[K]T, key K) bool {
	return ContainsInMap(m, key)
}

func (g *GenericUtility[T, K]) Marshal(model *T) ([]byte, error) {
	return Marshal(model)
}

func (g *GenericUtility[T, K]) Unmarshal(data []byte, model *T) (*T, error) {
	return Unmarshal(data, model)
}
