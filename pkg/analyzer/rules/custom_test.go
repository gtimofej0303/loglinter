package rules

import (
    "go/token"
    "testing"

    "golang.org/x/tools/go/analysis"
)

func TestCheckCustom(t *testing.T) {
    tests := []struct {
        name     string
        msg      string
        patterns []string
        words    []string
        wantErr  bool
    }{
        {"no rules", "user logged in", nil, nil, false},
        {"forbidden word match", "internal error occurred", nil, []string{"internal"}, true},
        {"forbidden word no match", "user logged in", nil, []string{"internal"}, false},
        {"pattern match", "credit card used", []string{"credit.?card"}, nil, true},
        {"pattern no match", "user logged in", []string{"credit.?card"}, nil, false},
        {"case insensitive word", "INTERNAL error", nil, []string{"internal"}, true},
        {"case insensitive pattern", "CreditCard", []string{"credit.?card"}, nil, true},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            reported := false
            pass := &analysis.Pass{
                Report: func(d analysis.Diagnostic) { reported = true },
            }
            CheckCustom(pass, tc.msg, token.NoPos, tc.patterns)
            if reported != tc.wantErr {
                t.Errorf("CheckCustom(%q) reported=%v, want %v", tc.msg, reported, tc.wantErr)
            }
        })
    }
}