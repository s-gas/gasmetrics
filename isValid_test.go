package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", false},
		{"float", "1234.56", false},
		{"negative", "-12", false},
		{"zero", "0", false},
		{"positive integer", "1", true},
	}

	for _, tt := range tests {
		if got := isValid(tt.input); got != tt.want {
			t.Errorf("isValid(%q) = %v; want: %v", tt.input, got, tt.want)
		}
	}
}
