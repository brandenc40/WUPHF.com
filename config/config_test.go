package config

import "testing"

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Loads with no errors", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadConfig(); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
