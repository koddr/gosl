package gosl

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

var resultRenders string

func BenchmarkRenderStyled(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = RenderStyled("Hello, World!", lipgloss.NewStyle().Foreground(lipgloss.Color("42")))
	}
	resultRenders = r
}

func TestRenderStyled(t *testing.T) {
	r := RenderStyled("", lipgloss.Style{})
	assert.EqualValues(t, "", r)

	r = RenderStyled("Hello, World!", lipgloss.NewStyle().Foreground(lipgloss.Color("42")))
	assert.EqualValues(t, "Hello, World!", r)

	g := Utility{} // tests for method

	r = g.RenderStyled("", lipgloss.Style{})
	assert.EqualValues(t, "", r)

	r = g.RenderStyled("Hello, World!", lipgloss.NewStyle().Foreground(lipgloss.Color("42")))
	assert.EqualValues(t, "Hello, World!", r)
}
