package rules

import (
	"go/token"
	"testing"

	"golang.org/x/tools/go/analysis"
)

func TestCheckEnglish(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"valid english only", "user logged in", false},
		{"valid with numbers", "4 users logged in", false},
		{"valid with punctuation", "connection failed, retrying", false},
		{"invalid russian letters", "пользователь вошёл", true},
		{"invalid mixed lang", "user вошёл", true},
		{"invalid chinese", "用户登录", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reported := false
			pass := &analysis.Pass{
				Report: func(d analysis.Diagnostic) {
					reported = true
				},
			}
			CheckEnglish(pass, tc.msg, token.NoPos)
			if reported != tc.wantErr {
				t.Errorf("checkEnglish(%q): reported=%v, want %v", tc.msg, reported, tc.wantErr)
			}
		})
	}
}