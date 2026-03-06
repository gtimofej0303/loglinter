package rules

import (
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Список чувствительных ключевых слов
var sensitiveKeywords = []string{
	"password",
	"pass",
	"passwd",
	"token",
	"apikey",
	"api_key",
	"secret",
	"private_key",
	"auth",
}

func CheckSensitive(pass *analysis.Pass, msg string, pos token.Pos) {
	lower := strings.ToLower(msg)
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lower, keyword) {
			pass.Reportf(pos, "log message must not contain sensitive data (%q): %q", keyword, msg)
			return
		}
	}
}
