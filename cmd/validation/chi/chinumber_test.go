package chi

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		args   args
		expErr error
	}{
		{
			name: "chi length too short",
			args: args{
				id: "123",
			},
			expErr: ErrChiLength,
		},
		{
			name: "chi length too long",
			args: args{
				id: "1234567891234567",
			},
			expErr: ErrChiLength,
		},
		{
			name: "chi contains non digits",
			args: args{
				id: "hello hi12",
			},
			expErr: ErrChiNonDigits,
		},
		{
			name: "valid chi",
			args: args{
				id: "1904851231",
			},
			expErr: nil,
		},
		{
			name: "zero day fails",
			args: args{
				id: "0011201234",
			},
			expErr: ErrChiInvalidDate,
		},
		{
			name: "zero month fails",
			args: args{
				id: "1100201234",
			},
			expErr: ErrChiInvalidDate,
		},
		{
			name: "non-leap year 29/02 fails",
			args: args{
				id: "2902191230",
			},
			expErr: ErrChiInvalidDate,
		},
		{
			name: "leap year check passes",
			args: args{
				id: "2902201230",
			},
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.args.id); !errors.Is(err, tt.expErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.expErr)
			}
		})
	}
}
