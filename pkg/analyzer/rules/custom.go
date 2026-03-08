package rules

import (
	"go/token"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

func CheckCustom(pass *analysis.Pass, msg string, pos token.Pos, patterns []string) {
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