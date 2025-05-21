package provider

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrimString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		prefixes []string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			length:   5,
			prefixes: []string{"he", "lo"},
			expected: "",
		},
		{
			name:     "length less than zero",
			input:    "hello world",
			length:   -1,
			prefixes: []string{"he", "lo"},
			expected: "",
		},
		{
			name:     "length zero",
			input:    "hello world",
			length:   0,
			prefixes: []string{"he", "lo"},
			expected: "",
		},
		{
			name:     "basic test with suffix",
			input:    "hello world-test",
			length:   5,
			prefixes: []string{"test"},
			expected: "hello",
		},
		{
			name:     "basic test with prefix",
			input:    "test-hello world",
			length:   5,
			prefixes: []string{"test"},
			expected: "hello",
		},
		{
			name:     "remove substrings",
			input:    "ead-staffapps-supplemental-appt-prod",
			length:   32,
			prefixes: []string{"studentapps", "staffapps"},
			expected: "ead-supplemental-appt-prod",
		},
		{
			name:     "remove substrings and trim to length",
			input:    "ead-staffapps-building-functionality-survey-test",
			length:   32,
			prefixes: []string{"studentapps", "staffapps"},
			expected: "ead-building-functionality-test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimString(tt.input, tt.length, tt.prefixes, true)
			require.Equal(t, tt.expected, result)
		})
	}
}
