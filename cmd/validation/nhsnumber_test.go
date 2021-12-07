package validation

import "testing"

func TestNhsNumberValidator(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "invalid length",
			id:      "123",
			wantErr: true,
		},
		{
			name:    "random string of length 10",
			id:      "Abc124398d",
			wantErr: true,
		},
		{
			name:    "invalid id",
			id:      "1234567890",
			wantErr: true,
		},
		{
			name:    "valid nhs number",
			id:      "2983396339",
			wantErr: false,
		},
		{
			name:    "valid nhs number 2",
			id:      "9000000025",
			wantErr: false,
		},
		{
			name:    "valid nhs number 3",
			id:      "9111231130",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NhsNumberValidator(tt.id); (err != nil) != tt.wantErr {
				t.Errorf("NhsNumberValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
