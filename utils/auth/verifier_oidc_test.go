package auth

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_verifyRoles(t *testing.T) {
	tests := []struct {
		actual   []string
		expected []string
		result   bool
	}{
		{
			actual:   []string{},
			expected: []string{},
			result:   true,
		},
		{
			actual:   []string{"additional"},
			expected: []string{},
			result:   true,
		},
		{
			actual:   []string{"matching"},
			expected: []string{"matching"},
			result:   true,
		},
		{
			actual:   []string{},
			expected: []string{"missing"},
			result:   false,
		},
		{
			actual:   []string{"mismatch-1"},
			expected: []string{"mismatch-2"},
			result:   false,
		},
	}

	for i, subtest := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, subtest.result, verifyRoles(subtest.actual, subtest.expected))
		})
	}
}
