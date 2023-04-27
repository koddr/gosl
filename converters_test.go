package gosl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var resultConvertersString string
var resultConvertersByte []byte

func BenchmarkToString_HelloWorld(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r, _ = ToString([]byte(`hello, world`))
	}
	resultConvertersString = r
}

func BenchmarkToBytes_HelloWorld(b *testing.B) {
	var r []byte
	for i := 0; i < b.N; i++ {
		r, _ = ToBytes("hello, world")
	}
	resultConvertersByte = r
}

func TestToString(t *testing.T) {
	s, err := ToString(nil)
	require.Error(t, err)

	s, err = ToString([]byte(`hello, world`))
	require.NoError(t, err)
	assert.EqualValues(t, s, "hello, world", "should be equal")
	assert.NotEqual(t, s, "wrong", "should not be equal")

	g := Utility{} // tests for method

	_, err = g.ToString(nil)
	require.Error(t, err)

	s, err = g.ToString([]byte(`hello, world`))
	require.NoError(t, err)
	assert.EqualValues(t, s, "hello, world", "should be equal")
	assert.NotEqual(t, s, "wrong", "should not be equal")
}

func TestToBytes(t *testing.T) {
	b, err := ToBytes("")
	require.Error(t, err)

	b, err = ToBytes("hello, world")
	require.NoError(t, err)
	assert.EqualValues(t, b, []byte(`hello, world`), "should be equal")
	assert.NotEqual(t, b, []byte(`wrong`), "should not be equal")

	g := Utility{} // tests for method

	_, err = g.ToBytes("")
	require.Error(t, err)

	b, err = g.ToBytes("hello, world")
	require.NoError(t, err)
	assert.EqualValues(t, b, []byte(`hello, world`), "should be equal")
	assert.NotEqual(t, b, []byte(`wrong`), "should not be equal")
}
