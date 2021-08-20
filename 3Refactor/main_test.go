package main

import (
	"testing"
)

func Test_findFirstStringInBracket(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "case 1",
			input:    "bi(sa)bit",
			expected: "sa",
		},
		{
			name:     "case 2",
			input:    "(bisa) bibit",
			expected: "bisa",
		},
		{
			name:     "case 3",
			input:    "bisa (bibit)",
			expected: "bibit",
		},
		{
			name:     "case 4",
			input:    "bi()bit",
			expected: "",
		},
		{
			name:     "case 5",
			input:    "()bibit",
			expected: "",
		},
		{
			name:     "case 6",
			input:    "bibit()",
			expected: "",
		},
		{
			name:     "case 7",
			input:    "bi)sa(bit",
			expected: "",
		},
		{
			name:     "case 8",
			input:    "bi(bit",
			expected: "",
		},
		{
			name:     "case 9",
			input:    "bi)bit",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := findFirstStringInBracket(tt.input); res != tt.expected {
				t.Errorf("unexpected result: %v. expected: %v", res, tt.expected)
			}
		})
	}
}
