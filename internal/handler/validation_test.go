package handler

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input1  string
		input2  int
		wantVal int
		wantErr bool
	}{
		{"empty", "", 100, 0, true},
		{"float", "1234.56", 100, 0, true},
		{"negative", "-12", 100, 0, true},
		{"zero", "0", 100, 0, true},
		{"positive integer", "1", 100, 1, false},
		{"smaller than last entry", "1", 3, 0, true},
	}

	for _, tt := range tests {
		gotVal, gotErr := validate(tt.input1, tt.input2)
		if (gotErr != nil) != tt.wantErr {
			t.Errorf("validate(%q, %v) = error = %v; wantErr: %v", tt.input1, tt.input2, gotErr, tt.wantErr)
		}
		if gotVal != tt.wantVal {
			t.Errorf("validate(%q, %v) = %v; want: %v", tt.input1, tt.input2, gotVal, tt.wantVal)
		}
	}
}
