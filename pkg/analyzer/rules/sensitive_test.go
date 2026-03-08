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
		extra []string
	}{
		{"valid clean message", "user logged in", false, nil},
		{"valid db connection", "database connection established", false, nil},
		{"invalid password keyword", "user password reset", true, nil},
		{"invalid token keyword", "received token from client", true, nil},
		{"invalid secret keyword", "loading secret config", true, []string{"secret"}},
		{"invalid api_key keyword", "api_key not found", true, nil},
		{"invalid apikey keyword", "apikey validation failed", true, nil},
		{"invalid auth keyword", "auth failed for user", true, nil},
		{"invalid private_key", "private_key loaded", true, nil},
		{"invalid passwd", "passwd is empty", true, nil},
		{"invalid pass keyword", "waiting pass validation", true, nil},
		{"invalid uppercase keyword", "User PASSWORD reset", true, nil},
		{"valid word passport", "passport validation", false, nil}, 
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reported := false
			pass := &analysis.Pass{
				Report: func(d analysis.Diagnostic) {
					reported = true
				},
			}
			CheckSensitiveWithExtra(pass, tc.msg, token.NoPos, tc.extra)
			if reported != tc.wantErr {
				t.Errorf("checkSensitive(%q): reported=%v, want %v", tc.msg, reported, tc.wantErr)
			}
		})
	}
}
