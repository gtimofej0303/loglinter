package rules

import (
	"go/token"
	"testing"

	"golang.org/x/tools/go/analysis"
)

func TestCheckLowercase(t *testing.T) {
	tests := []struct {
		name 	string
		msg 	string
		wantErr bool
	}{
		{"valid lowercase", "user logged in", false},
		{"valid starts with digit", "404 not found", false},
		{"valid empty string", "", false},
		{"invalid uppercase first letter", "User logged in", true},
		{"invalid all caps", "ERROR", true},
		{"invalid single uppercase", "A", true},
	}

	for _, tc := range tests{
		t.Run(tc.name, func(t *testing.T){
			reported := false
			pass := &analysis.Pass{
				Report: func(d analysis.Diagnostic) {
					reported = true
				},
			}
			CheckLowercase(pass, tc.msg, token.NoPos)
			if reported != tc.wantErr{
				t.Errorf("checkLowercase(%q): reported=%v, want %v", tc.msg, reported, tc.wantErr)
			}
		})
	}
}
