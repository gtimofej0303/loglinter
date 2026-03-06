package rules

import (
	"go/token"
	"testing"

	"golang.org/x/tools/go/analysis"
)

func TestCheckSensitive(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"valid clean message", "user logged in", false},
		{"valid db connection", "database connection established", false},
		{"invalid password keyword", "user password reset", true},
		{"invalid token keyword", "received token from client", true},
		{"invalid secret keyword", "loading secret config", true},
		{"invalid api_key keyword", "api_key not found", true},
		{"invalid apikey keyword", "apikey validation failed", true},
		{"invalid auth keyword", "auth failed for user", true},
		{"invalid private_key", "private_key loaded", true},
		{"invalid passwd", "passwd is empty", true},
		{"invalid pass keyword", "pass validation", true},
		{"invalid uppercase keyword", "User PASSWORD reset", true},
		{"valid word passport", "passport validation", true}, 
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reported := false
			pass := &analysis.Pass{
				Report: func(d analysis.Diagnostic) {
					reported = true
				},
			}
			CheckSensitive(pass, tc.msg, token.NoPos)
			if reported != tc.wantErr {
				t.Errorf("checkSensitive(%q): reported=%v, want %v", tc.msg, reported, tc.wantErr)
			}
		})
	}
}
