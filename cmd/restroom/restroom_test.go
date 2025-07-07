package main

import (
	"os"
	"testing"
)

func TestGetEnvOrDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
		envValue     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Missing value with default", args{"ShOUlDnt_Xist883322", "test899764", "nil"}, "test899764"},
		{"Missing value without default", args{"ShOUlDnt_Xist8321", "", "nil"}, ""},
		{"Blank value with default", args{"ShOUlDnt_Xist8320", "thedefault", ""}, "thedefault"},
		{"Blank value without default", args{"ShOUlDnt_Xist8319", "", ""}, ""},
		{"Value with default", args{"ShOUlDnt_Xist9318", "test3347", "test"}, "test"},
		{"Empty key with default", args{"", "testdefault", "nil"}, "testdefault"},
	}
	// INFO: "Blank value with default" tests that a blank value is replaced by the default value. Change test if behavior is changed.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.envValue != "nil" {
				err := os.Setenv(tt.args.key, tt.args.envValue)
				if err != nil {
					t.Errorf("Error setting environment variable (%v=%v): %v", tt.args.key, tt.args.envValue, err)
				}
			}
			if output := GetEnvOrDefault(tt.args.key, tt.args.defaultValue); output != tt.want {
				t.Errorf("GetEnvOrDefault(%v, %v) = %v, want %v", tt.args.key, tt.args.defaultValue, output, tt.want)
			}
		})
	}
}

// func TestGetEnvSettings(t *testing.T) {
// 	type args struct {
// 		envValues map[string]string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *EnvSettings
// 	}{
// 		{"All defaults", args{map[string]string{}}, &EnvSettings{"warn", "error", "", "9988", false, "/etc/restroom/auth.yaml"}},
// 		{"All values", args{map[string]string{
// 			"RESTROOM_LOG_LEVEL":    "info",
// 			"RESTROOM_STDERR_LEVEL": "debug",
// 			"RESTROOM_IP":           "