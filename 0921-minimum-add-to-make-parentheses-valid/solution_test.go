package _921_minimum_add_to_make_parentheses_valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solution(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"All Valid", "abc(de)f", 0},
		{"Mismatched Start", "a)b(c)d", 1},
		{"Mismatched End", "(a(b)c", 1},
		{"Empty String", "", 0},
		{"Continous Brackets", "(((((", 5},
		{"Complex Case", "a(b(c)d)e)()", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MinAddToMakeValid(tt.input))
		})
	}
}
