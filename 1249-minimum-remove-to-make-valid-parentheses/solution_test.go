package _249_minimum_remove_to_make_valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solution(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"All Valid", "abc(de)f", "abc(de)f"},
		{"Mismatched Start", "a)b(c)d", "ab(c)d"},
		{"Mismatched End", "(a(b)c", "(ab)c"},
		{"Empty String", "", ""},
		{"Complex Case", "a(b(c)d)e)()", "a(b(c)d)e()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MinRemoveToMakeValid(tt.input))
		})
	}
}
