package rules

import (
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

const forbidden = "?!#@$%^&*."

func CheckSpecialChars(pass *analysis.Pass, msg string, pos token.Pos) {
    for _, r := range msg {
        if unicode.Is(unicode.So, r) || unicode.Is(unicode.Sm, r) {
            pass.Reportf(pos, "log message must not contain special characters or emoji: %q", msg)
            return
        }

        if strings.ContainsRune(forbidden, r){
            pass.Reportf(pos, "log message must not contain special characters: %q", msg)
            return
        }
    }
}

func ContainsSpecialChars(msg string) bool {
	for _, r := range msg {
		if unicode.Is(unicode.So, r) || unicode.Is(unicode.Sm, r) {
			return true
		}
		if strings.ContainsRune(forbidden, r) {
			return true
		}
	}
	return false
}