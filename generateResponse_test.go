package main

import "testing"

func TestGenerateResponse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", "Invalid input: enter a positive integer"},
		{"negative", "-1", "Invalid input: enter a positive integer"},
		{"positive", "1", "Input received"},
		{"not a number", "hello", "Invalid input: enter a positive integer"},
	}
	for _, tt := range tests {
		if got := generateResponse(tt.input); got != tt.want {
			t.Errorf("generateResponse(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
