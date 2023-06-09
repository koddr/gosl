package gosl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equals("hello", "hello")
	}
}

func BenchmarkNotEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotEquals(42, 64)
	}
}

func TestEquals(t *testing.T) {
	b := Equals("hello", "hello")
	assert.EqualValues(t, b, true)

	b = Equals(42, 64)
	assert.EqualValues(t, b, false)

	g1 := GenericUtility[any, string]{} // tests for method

	b = g1.Equals("hello", "hello")
	assert.EqualValues(t, b, true)

	g2 := GenericUtility[any, int]{} // tests for method

	b = g2.Equals(42, 64)
	assert.EqualValues(t, b, false)
}

func TestNotEquals(t *testing.T) {
	b := NotEquals("hello", "hello")
	assert.EqualValues(t, b, false)

	b = NotEquals(42, 64)
	assert.EqualValues(t, b, true)

	g1 := GenericUtility[any, string]{} // tests for method

	b = g1.NotEquals("hello", "hello")
	assert.EqualValues(t, b, false)

	g2 := GenericUtility[any, int]{} // tests for method

	b = g2.NotEquals(42, 64)
	assert.EqualValues(t, b, true)
}
