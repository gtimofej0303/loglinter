package rules

import (
	"go/token"
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func CheckCustom(pass *analysis.Pass, msg string, pos token.Pos, patterns []string, words []string) {
	lower := strings.ToLower(msg)

	// Проверка запрещённых слов
	for _, word := range words {
		if strings.Contains(lower, strings.ToLower(word)) {
			pass.Reportf(pos, "log message must not contain forbidden word %q: %q", word, msg)
			return
		}
	}

	// Проверка regexp-паттернов
	for _, pattern := range patterns {
		re, err := regexp.Compile("(?i)" + pattern)
		if err != nil {
			continue
		}
		if re.MatchString(msg) {
			pass.Reportf(pos, "log message matches forbidden pattern %q: %q", pattern, msg)
			return
		}
	}
}