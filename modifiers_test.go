package gosl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkModifyByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := map[string]any{"order": map[string]any{"tags": "new"}}
		foundValue := "new"
		newValue := "paid"
		_, _ = ModifyByValue(m, foundValue, newValue)
	}
}

func TestModifyByValue(t *testing.T) {
	foundValue := "new"
	newValue := "paid"

	isFound, result := ModifyByValue(nil, foundValue, newValue)
	assert.EqualValues(t, isFound, false)

	m := map[string]any{"order": map[string]any{"tags": "new"}}
	modified := map[string]any{"order": map[string]any{"tags": "paid"}}

	isFound, result = ModifyByValue(m, foundValue, newValue)
	assert.EqualValues(t, isFound, true)
	assert.EqualValues(t, result, modified)

	u := Utility{} // tests for method

	isFound2, result2 := u.ModifyByValue(nil, foundValue, newValue)
	assert.EqualValues(t, isFound2, false)

	m2 := map[string]any{"order": map[string]any{"tags": "new"}}
	modified2 := map[string]any{"order": map[string]any{"tags": "paid"}}

	isFound2, result2 = u.ModifyByValue(m2, foundValue, newValue)
	assert.EqualValues(t, isFound2, true)
	assert.EqualValues(t, result2, modified2)
}
