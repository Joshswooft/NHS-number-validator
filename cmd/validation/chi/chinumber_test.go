package chi

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Joshswooft/nhs/cmd/validation/utils"
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

func TestGetGender(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    utils.Gender
		wantErr error
	}{
		{
			name: "invalid length fails",
			args: args{
				id: "1",
			},
			want:    utils.Female,
			wantErr: ErrChiLength,
		},
		{
			name: "non-digit characters fails",
			args: args{
				id: "12345678a0",
			},
			want:    utils.Female,
			wantErr: ErrChiNonDigits,
		},
		{
			name: "get male",
			args: args{
				id: "1904851231",
			},
			want:    utils.Male,
			wantErr: nil,
		},
		{
			name: "get female",
			args: args{
				id: "0123456789",
			},
			want:    utils.Female,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGender(tt.args.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetGender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDateOfBirth(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr error
	}{
		{
			name: "invalid length fails",
			args: args{
				id: "1",
			},
			wantErr: ErrChiLength,
		},
		{
			name: "non-digit characters fails",
			args: args{
				id: "12345678a0",
			},
			wantErr: ErrChiNonDigits,
		},
		{
			name: "gets date of birth",
			args: args{
				id: "1904851231",
			},
			want:    time.Date(1985, 04, 19, 0, 0, 0, 0, time.UTC),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDateOfBirth(tt.args.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetDateOfBirth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDateOfBirth() = %v, want %v", got, tt.want)
			}
		})
	}
}
