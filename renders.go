package gosl

import "github.com/charmbracelet/lipgloss"

// RenderStyled render a styled string with a given lipgloss.Style template
// using "charmbracelet/lipgloss" package.
//
// If s have an "" (empty) value returns zero-value for a string.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//		"log"
//
//		"github.com/charmbracelet/lipgloss"
//		"github.com/koddr/gosl"
//	)
//
//	func main() {
//		t := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Margin(1)
//
//		s := gosl.RenderStyled("This is a styled text", t)
//
//		fmt.Println(s)
//	}
func RenderStyled(str string, template lipgloss.Style) string {
	return template.Render(str)
}
