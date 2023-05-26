package gosl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var resultConcatString string

func BenchmarkConcat_String2(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = Concat("Lorem ipsum ", "dolor sit amet")
	}
	resultConcatString = r
}

func BenchmarkConcat_String8(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = Concat(
			"Lorem ipsum ",
			"dolor sit amet, ",
			"consectetur adipiscing elit, ",
			"sed do eiusmod ",
			"tempor incididunt ",
			"ut labore et ",
			"dolore magna ",
			"aliqua.",
		)
	}
	resultConcatString = r
}

func BenchmarkConcat_String32(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = Concat(
			"Lorem ipsum ",
			"dolor sit amet, ",
			"consectetur adipiscing elit, ",
			"sed do eiusmod ",
			"tempor incididunt ",
			"ut labore et ",
			"dolore magna ",
			"aliqua.",
			"Ut enim ",
			"ad minim veniam, ",
			"quis nostrud ",
			"exercitation ullamco ",
			"laboris nisi ",
			"ut aliquip ",
			"ex ",
			"ea ",
			"commodo consequat. ",
			"Duis aute ",
			"irure dolor ",
			"in reprehenderit ",
			"in voluptate velit ",
			"esse cillum ",
			"dolore eu ",
			"fugiat ",
			"nulla pariatur. ",
			"Excepteur sint ",
			"occaecat cupidatat ",
			"non proident, ",
			"sunt in culpa qui ",
			"officia ",
			"deserunt mollit ",
			"anim id ",
			"est laborum.",
		)
	}
	resultConcatString = r
}

func TestConcat(t *testing.T) {
	s := Concat()
	assert.EqualValues(t, s, "", "should be equal")

	s = Concat("Lorem ipsum ", "dolor ", "sit amet, ", "consectetur ", "adipiscing elit")
	assert.EqualValues(t, s, "Lorem ipsum dolor sit amet, consectetur adipiscing elit", "should be equal")
	assert.NotEqual(t, s, "wrong", "should not be equal")

	g := Utility{} // tests for method

	s = g.Concat()
	assert.EqualValues(t, s, "", "should be equal")

	s = g.Concat("Lorem ipsum ", "dolor ", "sit amet, ", "consectetur ", "adipiscing elit")
	assert.EqualValues(t, s, "Lorem ipsum dolor sit amet, consectetur adipiscing elit", "should be equal")
	assert.NotEqual(t, s, "wrong", "should not be equal")
}
