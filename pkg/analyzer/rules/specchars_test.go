package rules

import (
	"go/token"
	"testing"

	"golang.org/x/tools/go/analysis"
)

func TestCheckSpecialChars(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"valid clean message", "user logged in", false},
		{"valid with comma", "connected, loading data", false},
		{"valid with dash", "request timed-out", false},
		{"valid with colon", "error: connection refused", false},
		{"invalid exclamation mark", "success!", true},
		{"invalid question mark", "what happened?", true},
		{"invalid hash", "route #main", true},
		{"invalid at sign", "user @admin", true},
		{"invalid dollar sign", "cost $100", true},
		{"invalid percent sign", "loaded 100%", true},
		{"invalid caret", "value^2", true},
		{"invalid ampersand", "hi &&", true},
		{"invalid asterisk", "2 * 2", true},
		{"invalid emoji", "🎉", true},
		{"invalid math symbol", "∑", true},
		{"invalid dot", "something went wrong...", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reported := false
			pass := &analysis.Pass{
				Report: func(d analysis.Diagnostic) {
					reported = true
				},
			}
			CheckSpecialChars(pass, tc.msg, token.NoPos)
			if reported != tc.wantErr {
				t.Errorf("checkSpecialChars(%q): reported=%v, want %v", tc.msg, reported, tc.wantErr)
			}
		})
	}
}
