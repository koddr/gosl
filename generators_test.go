package gosl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var resultGenerators string

func BenchmarkRandomString_Size1(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r, _ = RandomString(1)
	}
	resultGenerators = r
}

func BenchmarkRandomString_Size8(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r, _ = RandomString(8)
	}
	resultGenerators = r
}

func BenchmarkRandomString_Size64(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r, _ = RandomString(64)
	}
	resultGenerators = r
}

func BenchmarkRandomString_Size512(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r, _ = RandomString(512)
	}
	resultGenerators = r
}

func BenchmarkRandomString_Size4096(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r, _ = RandomString(4096)
	}
	resultGenerators = r
}

func TestRandomString(t *testing.T) {
	_, err := RandomString(-1)
	require.Error(t, err)

	_, err = RandomString(0)
	require.Error(t, err)

	_, err = RandomString(8)
	require.NoError(t, err)

	g := Utility{} // tests for method

	_, err = g.RandomString(-1)
	require.Error(t, err)

	_, err = g.RandomString(0)
	require.Error(t, err)

	_, err = g.RandomString(8)
	require.NoError(t, err)
}
