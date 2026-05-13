package telegram

import "testing"

func TestIsPositiveInteger(t *testing.T) {
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
		if got := isPositiveInteger(tt.input); got != tt.want {
			t.Errorf("isPositiveInteger(%q) = %v; want: %v", tt.input, got, tt.want)
		}
	}
}
