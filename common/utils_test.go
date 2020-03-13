package common

import "testing"

func TestValidateEmailFormat(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Valid email", args{email: "test@gmail.com"}, false},
		{"Invalid email", args{email: "testjw3.@.com"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEmailFormat(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmailFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
