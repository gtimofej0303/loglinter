package rules

import (
    "go/token"
    "unicode"

    "golang.org/x/tools/go/analysis"
)

func CheckSpecialChars(pass *analysis.Pass, msg string, pos token.Pos) {
    for _, r := range msg {
        if unicode.Is(unicode.So, r) || unicode.Is(unicode.Sm, r) {
            pass.Reportf(pos, "log message must not contain special characters or emoji: %q", msg)
            return
        }

        if (r == '!') || (r == '?') || (r == '#') || (r == '@') || (r == '$') || (r == '%') || (r == '^') || (r == '&') || (r == '*') {
            pass.Reportf(pos, "log message must not contain special characters: %q", msg)
            return
        }
    }
}