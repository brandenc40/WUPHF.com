package models

import "testing"

func TestNewPhoneNumber(t *testing.T) {
	type args struct {
		pn string
	}
	tests := []struct {
		name    string
		args    args
		want    PhoneNumber
		wantErr bool
	}{
		{"911 fails", args{pn: "911"}, "", true},
		{"Letters in the phone number", args{pn: "72983jkj21"}, "", true},
		{"Valid number", args{pn: "563.343.5557"}, PhoneNumber("+15633435557"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPhoneNumber(tt.args.pn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
