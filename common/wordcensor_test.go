package common

import "testing"

func TestContainsCurseWords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Contains a curse word", args{str: "Fuck you"}, true},
		{"Doesn't contain a curse word", args{str: "Test"}, false},
		{"Empty string", args{str: ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsCurseWords(tt.args.str); got != tt.want {
				t.Errorf("ContainsCurseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
