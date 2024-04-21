package _071_simplify_path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplifyPath(t *testing.T) {

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"test 1",
			"/home/",
			"/home",
		},
		{"test 2",
			"/home/../opt/.",
			"/opt",
		},
		{"test 3",
			"/home/users/abc/../../../opt",
			"/opt",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, SimplifyPath(tt.path))
	}
}
