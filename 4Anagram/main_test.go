package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected bool
	}{
		{
			name:     "is valid anagram",
			input:    []string{"anagram", "nagaram"},
			expected: true,
		},
		{
			name:     "is invalid anagram - length diff",
			input:    []string{"red", "green"},
			expected: false,
		},
		{
			name:     "is invalid anagram - char diff",
			input:    []string{"ab", "ac"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isValidAnagram(tt.input[0], tt.input[1])
			if res != tt.expected {
				t.Errorf("unexpected result: %v. expected %v", res, tt.expected)
			}
		})
	}
}
